package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

const version string = "0.5.1"

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

	startServer()
}

func printVersion() {
	fmt.Println("junos_exporter")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("Author(s): Daniel Czerwonk")
	fmt.Println("Metric exporter for switches and routers running JunOS")
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
	reg.MustRegister(&JunosCollector{})

	promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog:      log.NewErrorLogger(),
		ErrorHandling: promhttp.ContinueOnError}).ServeHTTP(w, r)
}
