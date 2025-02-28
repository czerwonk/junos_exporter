// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/czerwonk/junos_exporter/internal/config"
	"github.com/czerwonk/junos_exporter/pkg/connector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/codes"

	log "github.com/sirupsen/logrus"
)

const version string = "0.14.0"

var (
	showVersion                 = flag.Bool("version", false, "Print version information.")
	listenAddress               = flag.String("web.listen-address", ":9326", "Address on which to expose metrics and web interface.")
	metricsPath                 = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	sshHosts                    = flag.String("ssh.targets", "", "Hosts to scrape")
	sshUsername                 = flag.String("ssh.user", "junos_exporter", "Username to use when connecting to junos devices using ssh")
	sshKeyFile                  = flag.String("ssh.keyfile", "", "Public key file to use when connecting to junos devices using ssh")
	sshKeyPassphrase            = flag.String("ssh.keyPassphrase", "", "Passphrase to decrypt key file if it's encrypted")
	sshPassword                 = flag.String("ssh.password", "", "Password to use when connecting to junos devices using ssh")
	sshReconnectInterval        = flag.Duration("ssh.reconnect-interval", 30*time.Second, "Duration to wait before reconnecting to a device after connection got lost")
	sshKeepAliveInterval        = flag.Duration("ssh.keep-alive-interval", 10*time.Second, "Duration to wait between keep alive messages")
	sshKeepAliveTimeout         = flag.Duration("ssh.keep-alive-timeout", 15*time.Second, "Duration to wait for keep alive message response")
	sshExpireTimeout            = flag.Duration("ssh.expire-timeout", 15*time.Minute, "Duration after an connection is terminated when it is not used")
	debug                       = flag.Bool("debug", false, "Show verbose debug output in log")
	alarmEnabled                = flag.Bool("alarm.enabled", true, "Scrape Alarm metrics")
	bgpEnabled                  = flag.Bool("bgp.enabled", true, "Scrape BGP metrics")
	ospfEnabled                 = flag.Bool("ospf.enabled", true, "Scrape OSPFv3 metrics")
	isisEnabled                 = flag.Bool("isis.enabled", false, "Scrape ISIS metrics")
	l2circuitEnabled            = flag.Bool("l2circuit.enabled", false, "Scrape l2circuit metrics")
	l2vpnEnabled                = flag.Bool("l2vpn.enabled", false, "Scrape l2vpn metrics")
	natEnabled                  = flag.Bool("nat.enabled", false, "Scrape NAT metrics")
	nat2Enabled                 = flag.Bool("nat2.enabled", false, "Scrape NAT2 metrics")
	ldpEnabled                  = flag.Bool("ldp.enabled", true, "Scrape ldp metrics")
	routingEngineEnabled        = flag.Bool("routingengine.enabled", true, "Scrape Routing Engine metrics")
	routesEnabled               = flag.Bool("routes.enabled", true, "Scrape routing table metrics")
	environmentEnabled          = flag.Bool("environment.enabled", true, "Scrape environment metrics")
	firewallEnabled             = flag.Bool("firewall.enabled", true, "Scrape Firewall count metrics")
	interfacesEnabled           = flag.Bool("interfaces.enabled", true, "Scrape interface metrics")
	interfaceDiagnosticsEnabled = flag.Bool("ifdiag.enabled", true, "Scrape optical interface diagnostic metrics")
	ipsecEnabled                = flag.Bool("ipsec.enabled", false, "Scrape IPSec metrics")
	securityEnabled             = flag.Bool("security.enabled", false, "Scrape security metrics")
	securityPoliciesEnabled     = flag.Bool("security_policies.enabled", false, "Scrape security policy metrics")
	storageEnabled              = flag.Bool("storage.enabled", true, "Scrape system storage metrics")
	fpcEnabled                  = flag.Bool("fpc.enabled", true, "Scrape line card metrics")
	accountingEnabled           = flag.Bool("accounting.enabled", false, "Scrape accounting flow metrics")
	interfaceQueuesEnabled      = flag.Bool("queues.enabled", false, "Scrape interface queue metrics")
	rpkiEnabled                 = flag.Bool("rpki.enabled", false, "Scrape rpki metrics")
	satelliteEnabled            = flag.Bool("satellite.enabled", false, "Scrape metrics from satellite devices")
	systemEnabled               = flag.Bool("system.enabled", false, "Scrape system metrics")
	macEnabled                  = flag.Bool("mac.enabled", false, "Scrape MAC address table metrics")
	alarmFilter                 = flag.String("alarms.filter", "", "Regex to filter for alerts to ignore")
	configFile                  = flag.String("config.file", "", "Path to config file")
	dynamicIfaceLabels          = flag.Bool("dynamic-interface-labels", true, "Parse interface descriptions to get labels dynamically")
	interfaceDescriptionRegex   = flag.String("interface-description-regex", "", "give a regex to retrieve the interface description labels")
	lsEnabled                   = flag.Bool("logical-systems.enabled", false, "Enable logical systems support")
	powerEnabled                = flag.Bool("power.enabled", true, "Scrape power metrics")
	lacpEnabled                 = flag.Bool("lacp.enabled", false, "Scrape LACP metrics")
	bfdEnabled                  = flag.Bool("bfd.enabled", false, "Scrape BFD metrics")
	vpwsEnabled                 = flag.Bool("vpws.enabled", false, "Scrape EVPN VPWS metrics")
	mplsLSPEnabled              = flag.Bool("mpls_lsp.enabled", false, "Scrape MPLS LSP metrics")
	licenseEnabled              = flag.Bool("license.enabled", false, "Scrape license metrics")
	tlsEnabled                  = flag.Bool("tls.enabled", false, "Enables TLS")
	tlsCertChainPath            = flag.String("tls.cert-file", "", "Path to TLS cert file")
	tlsKeyPath                  = flag.String("tls.key-file", "", "Path to TLS key file")
	tracingEnabled              = flag.Bool("tracing.enabled", false, "Enables tracing using OpenTelemetry")
	tracingProvider             = flag.String("tracing.provider", "", "Sets the tracing provider (stdout or collector)")
	tracingCollectorEndpoint    = flag.String("tracing.collector.grpc-endpoint", "", "Sets the tracing provider (stdout or collector)")
	subscriberEnabled           = flag.Bool("subscriber.enabled", false, "Scrape subscribers detail")
	macsecEnabled               = flag.Bool("macsec.enabled", true, "Scrape MACSec metrics")
	arpEnabled                  = flag.Bool("arps.enabled", true, "Scrape ARP metrics")
	poeEnabled                  = flag.Bool("poe.enabled", true, "Scrape PoE metrics")
	krtEnabled                  = flag.Bool("krt.enabled", false, "Scrape KRT queue metrics")
	cfg                         *config.Config
	devices                     []*connector.Device
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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdownTracing, err := initTracing(ctx)
	if err != nil {
		log.Fatalf("could not initialize tracing: %v", err)
	}
	defer shutdownTracing()

	initChannels(ctx)

	startServer()
}

func initChannels(ctx context.Context) {
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
			case <-ctx.Done():
				shutdown()
			case <-term:
				shutdown()
			}
		}
	}()
}

func shutdown() {
	log.Infoln("Closing connections to devices")
	connManager.CloseAll()
	os.Exit(0)
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

	devices, err = devicesForConfig(c)
	if err != nil {
		return err
	}
	cfg = c

	connManager = connectionManager()

	return nil
}

func reinitialize() error {
	configMu.Lock()
	defer configMu.Unlock()

	if connManager != nil {
		connManager.CloseAll()
		connManager = nil
	}

	return initialize()
}

func loadConfig() (*config.Config, error) {
	if len(*configFile) == 0 {
		return loadConfigFromFlags(), nil
	}

	log.Infoln("Loading config from", *configFile)
	b, err := os.ReadFile(*configFile)
	if err != nil {
		return nil, err
	}

	return config.Load(bytes.NewReader(b), *dynamicIfaceLabels)
}

func loadConfigFromFlags() *config.Config {
	c := config.New()
	c.Targets = strings.Split(*sshHosts, ",")
	c.LSEnabled = *lsEnabled
	c.IfDescReStr = *interfaceDescriptionRegex

	f := &c.Features
	f.Alarm = *alarmEnabled
	f.BGP = *bgpEnabled
	f.Environment = *environmentEnabled
	f.Firewall = *firewallEnabled
	f.Interfaces = *interfacesEnabled
	f.InterfaceDiagnostic = *interfaceDiagnosticsEnabled
	f.InterfaceQueue = *interfaceQueuesEnabled
	f.IPSec = *ipsecEnabled
	f.Security = *securityEnabled
	f.SecurityPolicies = *securityPoliciesEnabled
	f.ISIS = *isisEnabled
	f.NAT = *natEnabled
	f.NAT2 = *nat2Enabled
	f.OSPF = *ospfEnabled
	f.LDP = *ldpEnabled
	f.L2Circuit = *l2circuitEnabled
	f.L2Vpn = *l2vpnEnabled
	f.Routes = *routesEnabled
	f.RoutingEngine = *routingEngineEnabled
	f.Accounting = *accountingEnabled
	f.FPC = *fpcEnabled
	f.RPKI = *rpkiEnabled
	f.Storage = *storageEnabled
	f.Satellite = *satelliteEnabled
	f.System = *systemEnabled
	f.Power = *powerEnabled
	f.MAC = *macEnabled
	f.LACP = *lacpEnabled
	f.BFD = *bfdEnabled
	f.VPWS = *vpwsEnabled
	f.MPLSLSP = *mplsLSPEnabled
	f.License = *licenseEnabled
	f.Subscriber = *subscriberEnabled
	f.MACSec = *macsecEnabled
	f.ARP = *arpEnabled
	f.Poe = *poeEnabled
	f.KRT = *krtEnabled
	return c
}

func connectionManager() *connector.SSHConnectionManager {
	opts := []connector.Option{
		connector.WithReconnectInterval(*sshReconnectInterval),
		connector.WithKeepAliveInterval(*sshKeepAliveInterval),
		connector.WithKeepAliveTimeout(*sshKeepAliveTimeout),
		connector.WithExpiredConnectionTimeout(*sshExpireTimeout),
	}

	return connector.NewConnectionManager(opts...)
}

func startServer() {
	log.Infof("Starting JunOS exporter (Version: %s)", version)
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
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

	log.Infof("Listening for %s on %s (TLS: %v)", *metricsPath, *listenAddress, *tlsEnabled)
	if *tlsEnabled {
		log.Fatal(http.ListenAndServeTLS(*listenAddress, *tlsCertChainPath, *tlsKeyPath, nil))
		return
	}

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

	ctx, span := tracer.Start(r.Context(), "HandleMetricsRequest")
	defer span.End()

	reg := prometheus.NewRegistry()

	devs, err := devicesForRequest(r)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	logicalSystem := r.URL.Query().Get("ls")
	if !cfg.LSEnabled && logicalSystem != "" {
		err := fmt.Errorf("Logical systems not enabled but the logical system '%s' in parameters", logicalSystem)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), 400)
		return
	}

	c := newJunosCollector(ctx, devs, logicalSystem)
	reg.MustRegister(c)

	l := log.New()
	l.Level = log.ErrorLevel

	promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog:      l,
		ErrorHandling: promhttp.ContinueOnError,
	}).ServeHTTP(w, r)
}

func devicesForRequest(r *http.Request) ([]*connector.Device, error) {
	reqTarget := r.URL.Query().Get("target")
	if reqTarget == "" {
		return devices, nil
	}

	for _, d := range devices {
		if d.Host == reqTarget {
			return []*connector.Device{d}, nil
		}
	}

	for _, dc := range cfg.Devices {
		if !dc.IsHostPattern {
			continue
		}

		if dc.HostPattern.MatchString(reqTarget) {
			d, err := deviceFromDeviceConfig(dc, reqTarget, cfg)
			if err != nil {
				return nil, err
			}

			return []*connector.Device{d}, nil
		}
	}

	return nil, fmt.Errorf("the target '%s' is not defined in the configuration file", reqTarget)
}
