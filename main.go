package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/czerwonk/junos_exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

const version string = "0.6.3"

var (
	showVersion          = flag.Bool("version", false, "Print version information.")
	listenAddress        = flag.String("web.listen-address", ":9326", "Address on which to expose metrics and web interface.")
	metricsPath          = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	sshHosts             = flag.String("ssh.targets", "", "Hosts to scrape")
	sshUsername          = flag.String("ssh.user", "junos_exporter", "Username to use when connecting to junos devices using ssh")
	sshKeyFile           = flag.String("ssh.keyfile", "junos_exporter", "Public key file to use when connecting to junos devices using ssh")
	debug                = flag.Bool("debug", false, "Show verbose debug output in log")
	bgpEnabled           = flag.Bool("bgp.enabled", true, "Scrape BGP metrics")
	ospfEnabled          = flag.Bool("ospf.enabled", true, "Scrape OSPFv3 metrics")
	isisEnabled          = flag.Bool("isis.enabled", false, "Scrape ISIS metrics")
	routingEngineEnabled = flag.Bool("routingengine.enabled", true, "Scrape Routing Engine metrics")
	routesEnabled        = flag.Bool("routes.enabled", true, "Scrape routing table metrics")
	environmentEnabled   = flag.Bool("environment.enabled", true, "Scrape environment metrics")
	ifDiagnEnabled       = flag.Bool("ifdiag.enabled", true, "Scrape optical interface diagnostic metrics")
	alarmFilter          = flag.String("alarms.filter", "", "Regex to filter for alerts to ignore")
	configFile           = flag.String("config.file", "", "Path to config file")
	cfg                  *config.Config
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
	f.BPG = *bgpEnabled
	f.Environment = *environmentEnabled
	f.InterfaceDiagnostic = *ifDiagnEnabled
	f.ISIS = *isisEnabled
	f.OSPF = *ospfEnabled
	f.Routes = *routesEnabled
	f.RoutingEngine = *routingEngineEnabled

	return c
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

	c := newJunosCollector()
	reg.MustRegister(c)

	promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog:      log.NewErrorLogger(),
		ErrorHandling: promhttp.ContinueOnError}).ServeHTTP(w, r)
}
