package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/czerwonk/junos_exporter/connector"
	"github.com/pkg/errors"

	"github.com/czerwonk/junos_exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

const version string = "0.9.2"

var (
	showVersion                 = flag.Bool("version", false, "Print version information.")
	ignoreConfigTargets         = flag.Bool("config.ignore-targets", false, "Ignore check if target is specified in config")
	listenAddress               = flag.String("web.listen-address", ":9326", "Address on which to expose metrics and web interface.")
	metricsPath                 = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	sshHosts                    = flag.String("ssh.targets", "", "Hosts to scrape")
	sshUsername                 = flag.String("ssh.user", "junos_exporter", "Username to use when connecting to junos devices using ssh")
	sshKeyFile                  = flag.String("ssh.keyfile", "", "Public key file to use when connecting to junos devices using ssh")
	sshPassword                 = flag.String("ssh.password", "", "Password to use when connecting to junos devices using ssh")
	sshReconnectInterval        = flag.Duration("ssh.reconnect-interval", 30*time.Second, "Duration to wait before reconnecting to a device after connection got lost")
	sshKeepAliveInterval        = flag.Duration("ssh.keep-alive-interval", 10*time.Second, "Duration to wait between keep alive messages")
	sshKeepAliveTimeout         = flag.Duration("ssh.keep-alive-timeout", 15*time.Second, "Duration to wait for keep alive message response")
	debug                       = flag.Bool("debug", false, "Show verbose debug output in log")
	bgpEnabled                  = flag.Bool("bgp.enabled", true, "Scrape BGP metrics")
	ospfEnabled                 = flag.Bool("ospf.enabled", true, "Scrape OSPFv3 metrics")
	isisEnabled                 = flag.Bool("isis.enabled", false, "Scrape ISIS metrics")
	l2circuitEnabled            = flag.Bool("l2circuit.enabled", false, "Scrape l2circuit metrics")
	natEnabled                  = flag.Bool("nat.enabled", false, "Scrape NAT metrics")
	ldpEnabled                  = flag.Bool("ldp.enabled", true, "Scrape ldp metrics")
	routingEngineEnabled        = flag.Bool("routingengine.enabled", true, "Scrape Routing Engine metrics")
	routesEnabled               = flag.Bool("routes.enabled", true, "Scrape routing table metrics")
	environmentEnabled          = flag.Bool("environment.enabled", true, "Scrape environment metrics")
	firewallEnabled             = flag.Bool("firewall.enabled", true, "Scrape Firewall count metrics")
	interfacesEnabled           = flag.Bool("interfaces.enabled", true, "Scrape interface metrics")
	interfaceDiagnosticsEnabled = flag.Bool("ifdiag.enabled", true, "Scrape optical interface diagnostic metrics")
	ipsecEnabled                = flag.Bool("ipsec.enabled", false, "Scrape IPSec metrics")
	storageEnabled              = flag.Bool("storage.enabled", true, "Scrape system storage metrics")
	accountingEnabled           = flag.Bool("accounting.enabled", false, "Scrape accounting flow metrics")
	alarmFilter                 = flag.String("alarms.filter", "", "Regex to filter for alerts to ignore")
	configFile                  = flag.String("config.file", "", "Path to config file")
	cfg                         *config.Config
	connManager                 *connector.SSHConnectionManager
	reloadCh                    chan chan error
	configMu                    sync.RWMutex
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

	err := initialize()
	if err != nil {
		log.Fatalf("could not initialize exporter. %v", err)
	}

	initChannels()

	startServer()
}

func initChannels() {
	hup := make(chan os.Signal, 1)
	signal.Notify(hup, syscall.SIGHUP)

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)

	reloadCh = make(chan chan error)
	go func() {
		for {
			select {
			case <-hup:
				log.Infoln("Reload signal received as SIGHUP")
				if err := reinitialize(); err != nil {
					log.Errorf("Error reloading config: %s", err)
				}
			case rc := <-reloadCh:
				log.Infoln("Reload signal received via POST")
				if err := reinitialize(); err != nil {
					log.Errorf("Error reloading config: %s", err)
					rc <- err
				} else {
					rc <- nil
				}
			case <-term:
				log.Infoln("Closing connections to devices")
				connManager.Close()
			}
		}
	}()
}

func printVersion() {
	fmt.Println("junos_exporter")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("Author(s): Daniel Czerwonk")
	fmt.Println("Metric exporter for switches and routers running JunOS")
}

func initialize() error {
	c, err := loadConfig()
	if err != nil {
		return err
	}
	cfg = c

	connManager, err = connectionManager()
	if err != nil {
		log.Fatalf("could initialize connection manager. %v", err)
	}

	return nil
}

func reinitialize() error {
	configMu.Lock()
	defer configMu.Unlock()

	if connManager != nil {
		connManager.Close()
		connManager = nil
	}

	return initialize()
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
	f.Firewall = *firewallEnabled
	f.Interfaces = *interfacesEnabled
	f.InterfaceDiagnostic = *interfaceDiagnosticsEnabled
	f.Ipsec = *ipsecEnabled
	f.ISIS = *isisEnabled
	f.NAT = *natEnabled
	f.OSPF = *ospfEnabled
	f.LDP = *ldpEnabled
	f.L2Circuit = *l2circuitEnabled
	f.Routes = *routesEnabled
	f.RoutingEngine = *routingEngineEnabled
	f.Accounting = *accountingEnabled

	return c
}

func connectionManager() (*connector.SSHConnectionManager, error) {
	opts := []connector.Option{
		connector.WithReconnectInterval(*sshReconnectInterval),
		connector.WithKeepAliveInterval(*sshKeepAliveInterval),
		connector.WithKeepAliveTimeout(*sshKeepAliveTimeout),
	}

	if *sshKeyFile != "" {
		return connectionManagerWithPublicKey(opts)
	}

	if *sshPassword != "" {
		return connector.NewConnectionManager(*sshUsername,
			connector.AuthByPassword(*sshPassword), opts...), nil
	}

	if cfg.Password != "" {
		return connector.NewConnectionManager(*sshUsername,
			connector.AuthByPassword(cfg.Password), opts...), nil
	}

	return nil, errors.New("no valid authentication method available")
}

func connectionManagerWithPublicKey(opts []connector.Option) (*connector.SSHConnectionManager, error) {
	f, err := os.Open(*sshKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "could not open ssh key file")
	}
	defer f.Close()

	auth, err := connector.AuthByKey(f)
	if err != nil {
		return nil, errors.Wrap(err, "could not load ssh private key file")
	}

	return connector.NewConnectionManager(*sshUsername, auth, opts...), err
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
	http.HandleFunc("/-/reload", updateConfiguration)

	log.Infof("Listening for %s on %s\n", *metricsPath, *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

func updateConfiguration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		rc := make(chan error)
		reloadCh <- rc
		if err := <-rc; err != nil {
			http.Error(w, fmt.Sprintf("failed to reload config: %s", err), http.StatusInternalServerError)
		}
	default:
		log.Errorf("POST method expected")
		http.Error(w, "POST method expected", 400)
	}
}

func handleMetricsRequest(w http.ResponseWriter, r *http.Request) {
	configMu.RLock()
	defer configMu.RUnlock()

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

	if *ignoreConfigTargets {
		return []string{reqTarget}, nil
	}

	return nil, fmt.Errorf("the target '%s' is not defined in the configuration file", reqTarget)
}
