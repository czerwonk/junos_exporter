package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"strings"

	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

const version string = "0.3.0"

var (
	showVersion   = flag.Bool("version", false, "Print version information.")
	listenAddress = flag.String("web.listen-address", ":9326", "Address on which to expose metrics and web interface.")
	metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	snmpTargets   = flag.String("snmp.targets", "127.0.0.1", "Addresses or hostnames of switches or routers (comma separated)")
	snmpCommunity = flag.String("snmp.community", "default", "Community allowed to access SNMP")
	mutex         = &sync.Mutex{}
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
	mutex.Lock()
	defer mutex.Unlock()

	reg := prometheus.NewRegistry()
	targets := strings.Split(*snmpTargets, ",")
	reg.MustRegister(NewJunosCollector(targets, *snmpCommunity))

	promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog:      log.NewErrorLogger(),
		ErrorHandling: promhttp.ContinueOnError}).ServeHTTP(w, r)
}
