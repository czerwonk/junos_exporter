package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/czerwonk/junos_exporter/connector"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"

	"github.com/czerwonk/junos_exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

const version string = "0.8.0"

var (
	showVersion                 = flag.Bool("version", false, "Print version information.")
	listenAddress               = flag.String("web.listen-address", ":9326", "Address on which to expose metrics and web interface.")
	metricsPath                 = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	sshHosts                    = flag.String("ssh.targets", "", "Hosts to scrape")
	sshUsername                 = flag.String("ssh.user", "junos_exporter", "Username to use when connecting to junos devices using ssh")
	sshKeyFile                  = flag.String("ssh.keyfile", "", "Public key file to use when connecting to junos devices using ssh")
	sshPassword                 = flag.String("ssh.password", "", "Password to use when connecting to junos devices using ssh")
	debug                       = flag.Bool("debug", false, "Show verbose debug output in log")
	bgpEnabled                  = flag.Bool("bgp.enabled", true, "Scrape BGP metrics")
	ospfEnabled                 = flag.Bool("ospf.enabled", true, "Scrape OSPFv3 metrics")
	isisEnabled                 = flag.Bool("isis.enabled", false, "Scrape ISIS metrics")
	l2circuitEnabled            = flag.Bool("l2circuit.enabled", false, "Scrape l2circuit metrics")
	routingEngineEnabled        = flag.Bool("routingengine.enabled", true, "Scrape Routing Engine metrics")
	routesEnabled               = flag.Bool("routes.enabled", true, "Scrape routing table metrics")
	environmentEnabled          = flag.Bool("environment.enabled", true, "Scrape environment metrics")
	interfacesEnabled           = flag.Bool("interfaces.enabled", true, "Scrape interface metrics")
	interfaceDiagnosticsEnabled = flag.Bool("ifdiag.enabled", true, "Scrape optical interface diagnostic metrics")
	storageEnabled              = flag.Bool("storage.enabled", true, "Scrape system storage metrics")
	alarmFilter                 = flag.String("alarms.filter", "", "Regex to filter for alerts to ignore")
	configFile                  = flag.String("config.file", "", "Path to config file")
	cfg                         *config.Config
	connManager                 *connector.SSHConnectionManager
)

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: junos_exporter [ ... ]\n\nParameters:")
		fmt.Println()
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}

	c, err := loadConfig()
	if err != nil {
		log.Fatalf("could not load config file. %v", err)
	}
	cfg = c

	connManager, err = connectionManager()
	if err != nil {
		log.Fatalf("could initialize connection manager. %v", err)
	}
	defer connManager.Close()

	startServer()
}

func printVersion() {
	fmt.Println("junos_exporter")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("Author(s): Daniel Czerwonk")
	fmt.Println("Metric exporter for switches and routers running JunOS")
}

func loadConfig() (*config.Config, error) {
	if len(*configFile) == 0 {
		return loadConfigFromFlags(), nil
	}

	log.Infoln("Loading config from", *configFile)
	b, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return nil, err
	}

	return config.Load(bytes.NewReader(b))
}

func loadConfigFromFlags() *config.Config {
	c := config.New()
	c.Targets = strings.Split(*sshHosts, ",")

	f := &c.Features
	f.BGP = *bgpEnabled
	f.Environment = *environmentEnabled
	f.Interfaces = *interfacesEnabled
	f.InterfaceDiagnostic = *interfaceDiagnosticsEnabled
	f.ISIS = *isisEnabled
	f.OSPF = *ospfEnabled
	f.L2Circuit = *l2circuitEnabled
	f.Routes = *routesEnabled
	f.RoutingEngine = *routingEngineEnabled

	return c
}

func loadPublicKeyFile(r io.Reader) (ssh.AuthMethod, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "could not read from reader")
	}

	key, err := ssh.ParsePrivateKey(b)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse private key")
	}

	return ssh.PublicKeys(key), nil
}

func connectionManager() (*connector.SSHConnectionManager, error) {
	if *sshKeyFile != "" {
		f, err := os.Open(*sshKeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "could not open ssh key file")
		}
		defer f.Close()

		pk, err := loadPublicKeyFile(f)
		auth := []ssh.AuthMethod{pk}
		if err != nil {
			return nil, errors.Wrap(err, "could not open ssh key file")
		}
		return connector.NewConnectionManager(*sshUsername, auth), err
	} else if *sshPassword != "" {
		auth := []ssh.AuthMethod{ssh.Password(*sshPassword)}
		return connector.NewConnectionManager(*sshUsername, auth), nil
	} else if cfg.Password != "" {
		auth := []ssh.AuthMethod{ssh.Password(cfg.Password)}
		return connector.NewConnectionManager(*sshUsername, auth), nil
	}
	return nil, errors.New("no valid authentication method available")
}

func startServer() {
	log.Infof("Starting JunOS exporter (Version: %s)\n", version)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>JunOS Exporter (Version ` + version + `)</title></head>
			<body>
			<h1>JunOS Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			<h2>More information:</h2>
			<p><a href="https://github.com/czerwonk/junos_exporter">github.com/czerwonk/junos_exporter</a></p>
			</body>
			</html>`))
	})
	http.HandleFunc(*metricsPath, handleMetricsRequest)

	log.Infof("Listening for %s on %s\n", *metricsPath, *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

func handleMetricsRequest(w http.ResponseWriter, r *http.Request) {
	reg := prometheus.NewRegistry()

	targets, err := targetsForRequest(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	c := newJunosCollector(targets, connManager)
	reg.MustRegister(c)

	promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog:      log.NewErrorLogger(),
		ErrorHandling: promhttp.ContinueOnError}).ServeHTTP(w, r)
}

func targetsForRequest(r *http.Request) ([]string, error) {
	reqTarget := r.URL.Query().Get("target")
	if reqTarget == "" {
		return cfg.Targets, nil
	}

	for _, t := range cfg.Targets {
		if t == reqTarget {
			return []string{t}, nil
		}
	}

	return nil, fmt.Errorf("the target '%s' is not defined in the configuration file", reqTarget)
}
