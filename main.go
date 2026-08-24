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
	"github.com/czerwonk/junos_exporter/internal/log/slogadapter"
	"github.com/czerwonk/junos_exporter/pkg/connector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/exporter-toolkit/web"
	"go.opentelemetry.io/otel/codes"

	log "github.com/sirupsen/logrus"
)

const version string = "0.16.2"

var (
	showVersion                 = flag.Bool("version", false, "Print version information.")
	listenAddress               = flag.String("web.listen-address", ":9326", "Address on which to expose metrics and web interface.")
	metricsPath                 = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	sshHosts                    = flag.String("ssh.targets", "", "Hosts to scrape")
	sshUsername                 = flag.String("ssh.user", "junos_exporter", "Username to use when connecting to junos devices using ssh")
	sshKeyFile                  = flag.String("ssh.keyfile", "", "Public key file to use when connecting to junos devices using ssh")
	sshKeyPassphrase            = flag.String("ssh.keyPassphrase", "", "Passphrase to decrypt key file if it's encrypted (mutually exclusive with -ssh.keyPassphraseEnv and -ssh.keyPassphraseFile)")
	sshKeyPassphraseEnv         = flag.String("ssh.keyPassphraseEnv", "", "Name of an environment variable to read the SSH key passphrase from")
	sshKeyPassphraseFile        = flag.String("ssh.keyPassphraseFile", "", "Path to a file containing the SSH key passphrase (trailing newline trimmed)")
	sshPassword                 = flag.String("ssh.password", "", "Password to use when connecting to junos devices using ssh (mutually exclusive with -ssh.passwordEnv and -ssh.passwordFile)")
	sshPasswordEnv              = flag.String("ssh.passwordEnv", "", "Name of an environment variable to read the SSH password from")
	sshPasswordFile             = flag.String("ssh.passwordFile", "", "Path to a file containing the SSH password (trailing newline trimmed)")
	sshReconnectInterval        = flag.Duration("ssh.reconnect-interval", 30*time.Second, "Duration to wait before reconnecting to a device after connection got lost")
	sshKeepAliveInterval        = flag.Duration("ssh.keep-alive-interval", 10*time.Second, "Duration to wait between keep alive messages")
	sshKeepAliveTimeout         = flag.Duration("ssh.keep-alive-timeout", 15*time.Second, "Duration to wait for keep alive message response")
	sshExpireTimeout            = flag.Duration("ssh.expire-timeout", 15*time.Minute, "Duration after an connection is terminated when it is not used")
	debug                       = flag.Bool("debug", false, "Show verbose debug output in log")
	alarmEnabled                = flag.Bool("alarm.enabled", true, "Scrape Alarm metrics")
	ntpEnabled                  = flag.Bool("ntp.enabled", false, "Scrape NTP metrics")
	bgpEnabled                  = flag.Bool("bgp.enabled", true, "Scrape BGP metrics")
	dot1xEnabled                = flag.Bool("dot1x.enabled", false, "Scrape dot1x metrics")
	ospfEnabled                 = flag.Bool("ospf.enabled", true, "Scrape OSPFv3 metrics")
	isisEnabled                 = flag.Bool("isis.enabled", true, "Scrape ISIS metrics")
	l2circuitEnabled            = flag.Bool("l2circuit.enabled", false, "Scrape l2circuit metrics")
	l2vpnEnabled                = flag.Bool("l2vpn.enabled", false, "Scrape l2vpn metrics")
	natEnabled                  = flag.Bool("nat.enabled", false, "Scrape NAT metrics")
	nat2Enabled                 = flag.Bool("nat2.enabled", false, "Scrape NAT2 metrics")
	ldpEnabled                  = flag.Bool("ldp.enabled", true, "Scrape ldp metrics")
	routingEngineEnabled        = flag.Bool("routingengine.enabled", true, "Scrape Routing Engine metrics")
	routesEnabled               = flag.Bool("routes.enabled", true, "Scrape routing table metrics")
	environmentEnabled          = flag.Bool("environment.enabled", true, "Scrape environment metrics")
	evpnEnabled                 = flag.Bool("evpn.enabled", false, "Scrape EVPN instance, detail tables (interfaces/IRBs/bridge-domains/ESIs), duplicate-MAC, and L3 context metrics")
	evpnIPPrefixEnabled         = flag.Bool("evpn_ip_prefix.enabled", false, "Scrape EVPN Type-5 (IP-prefix) database metrics; potentially large on busy fabrics")
	firewallEnabled             = flag.Bool("firewall.enabled", true, "Scrape Firewall count metrics")
	interfacesEnabled           = flag.Bool("interfaces.enabled", true, "Scrape interface metrics")
	interfaceDiagnosticsEnabled = flag.Bool("ifdiag.enabled", true, "Scrape optical interface diagnostic metrics")
	ipsecEnabled                = flag.Bool("ipsec.enabled", false, "Scrape IPSec metrics")
	securityEnabled             = flag.Bool("security.enabled", false, "Scrape security metrics")
	securityIKEEnabled          = flag.Bool("security_ike.enabled", false, "Scrape security IKE metrics")
	securityPoliciesEnabled     = flag.Bool("security_policies.enabled", false, "Scrape security policy metrics")
	storageEnabled              = flag.Bool("storage.enabled", false, "Scrape system storage metrics")
	fpcEnabled                  = flag.Bool("fpc.enabled", false, "Scrape line card metrics")
	accountingEnabled           = flag.Bool("accounting.enabled", false, "Scrape accounting flow metrics")
	interfaceQueuesEnabled      = flag.Bool("queues.enabled", true, "Scrape interface queue metrics")
	rpkiEnabled                 = flag.Bool("rpki.enabled", false, "Scrape rpki metrics")
	rpmEnabled                  = flag.Bool("rpm.enabled", false, "Scrape RPM metrics")
	satelliteEnabled            = flag.Bool("satellite.enabled", false, "Scrape metrics from satellite devices")
	systemEnabled               = flag.Bool("system.enabled", false, "Scrape system metrics")
	macEnabled                  = flag.Bool("mac.enabled", false, "Scrape MAC address table metrics")
	alarmFilter                 = flag.String("alarms.filter", "", "Regex to filter for alerts to ignore")
	firewallFilterNameRegex     = flag.String("firewall.filter-name-regex", "", "Regex to filter firewall filters by name")
	configFile                  = flag.String("config.file", "", "Path to config file")
	dynamicIfaceLabels          = flag.Bool("dynamic-interface-labels", true, "Parse interface descriptions to get labels dynamically")
	interfaceDescriptionRegex   = flag.String("interface-description-regex", "", "give a regex to retrieve the interface description labels")
	interfaceNameRegex          = flag.String("interfaces.name-regex", "", "Regex to filter interfaces by name")
	lsEnabled                   = flag.Bool("logical-systems.enabled", false, "Enable logical systems support")
	powerEnabled                = flag.Bool("power.enabled", false, "Scrape power metrics")
	lldpEnabled                 = flag.Bool("lldp.enabled", false, "Scrape LLDP metrics")
	lacpEnabled                 = flag.Bool("lacp.enabled", false, "Scrape LACP metrics")
	bfdEnabled                  = flag.Bool("bfd.enabled", false, "Scrape BFD metrics")
	clusterEnabled              = flag.Bool("cluster.enabled", false, "Scrape chassis cluster metrics")
	virtualChassisEnabled       = flag.Bool("virtual_chassis.enabled", false, "Scrape virtual chassis metrics")
	vrrpEnabled                 = flag.Bool("vrrp.enabled", false, "Scrape VRRP metrics")
	vpwsEnabled                 = flag.Bool("vpws.enabled", false, "Scrape EVPN VPWS metrics")
	mplsLSPEnabled              = flag.Bool("mpls_lsp.enabled", false, "Scrape MPLS LSP metrics")
	licenseEnabled              = flag.Bool("license.enabled", false, "Scrape license metrics")
	tlsEnabled                  = flag.Bool("tls.enabled", false, "Enables TLS")
	tlsCertChainPath            = flag.String("tls.cert-file", "", "Path to TLS cert file")
	tlsKeyPath                  = flag.String("tls.key-file", "", "Path to TLS key file")
	webConfigFile               = flag.String("web.config.file", "", "Path to web-config YAML (TLS + basic-auth, see prometheus/exporter-toolkit). When set, overrides -tls.* flags.")
	tracingEnabled              = flag.Bool("tracing.enabled", false, "Enables tracing using OpenTelemetry")
	tracingProvider             = flag.String("tracing.provider", "", "Sets the tracing provider (stdout or collector)")
	tracingCollectorEndpoint    = flag.String("tracing.collector.grpc-endpoint", "", "Sets the tracing provider (stdout or collector)")
	subscriberEnabled           = flag.Bool("subscriber.enabled", false, "Scrape subscribers detail")
	macsecEnabled               = flag.Bool("macsec.enabled", true, "Scrape MACSec metrics")
	arpEnabled                  = flag.Bool("arps.enabled", false, "Scrape ARP metrics")
	poeEnabled                  = flag.Bool("poe.enabled", false, "Scrape PoE metrics")
	ddosProtectionEnabled       = flag.Bool("ddos_protection.enabled", false, "Scrape DDoS protection metrics")
	krtEnabled                  = flag.Bool("krt.enabled", false, "Scrape KRT queue metrics")
	twampEnabled                = flag.Bool("twamp.enabled", false, "Scrape TWAMP metrics")
	systemstatisticsEnabled     = flag.Bool("systemstatistics.enabled", true, "Scrape system statistics metrics")
	ufdEnabled                  = flag.Bool("ufd.enabled", false, "Scrape UFD (uplink-failure-detection) metrics")
	mnhaEnabled                 = flag.Bool("mnha.enabled", false, "Scrape MNHA (Mixed/Multi-Node High Availability) metrics")
	mnhaSRGIDs                  = flag.String("mnha.srg-ids", "0", "Comma-separated list of MNHA services-redundancy-group IDs to scrape")
	securityNATEnabled          = flag.Bool("security_nat.enabled", false, "Scrape security NAT (source NAT pool usage and rule hit counters) metrics")
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

	if err := resolveSSHSecrets(); err != nil {
		log.Fatalf("could not resolve ssh credentials: %v", err)
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

	initChannels(ctx, cancel)

	go func() {
		if err := startServer(); err != nil {
			log.Errorf("server stopped unexpectedly: %v", err)
			cancel()
		}
	}()

	<-ctx.Done()
	log.Infoln("Closing connections to devices")
	connManager.CloseAll()
}

func initChannels(ctx context.Context, cancel context.CancelFunc) {
	hup := make(chan os.Signal, 1)
	signal.Notify(hup, syscall.SIGHUP)

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGTERM)

	reloadCh = make(chan chan error)
	go handleSignals(ctx, cancel, hup, term)
}

func handleSignals(ctx context.Context, cancel context.CancelFunc, hup, term <-chan os.Signal) {
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
			return
		case <-term:
			cancel()
			return
		}
	}
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

// resolveSSHSecrets materialises both *sshKeyPassphrase and *sshPassword from
// their respective literal flag / environment variable / file sources. Within
// each group, the three sources are mutually exclusive.
func resolveSSHSecrets() error {
	if err := resolveSecretFromSources(
		sshKeyPassphrase, sshKeyPassphraseEnv, sshKeyPassphraseFile,
		"-ssh.keyPassphrase", "-ssh.keyPassphraseEnv", "-ssh.keyPassphraseFile",
	); err != nil {
		return fmt.Errorf("ssh key passphrase: %w", err)
	}

	if err := resolveSecretFromSources(
		sshPassword, sshPasswordEnv, sshPasswordFile,
		"-ssh.password", "-ssh.passwordEnv", "-ssh.passwordFile",
	); err != nil {
		return fmt.Errorf("ssh password: %w", err)
	}

	return nil
}

// resolveSecretFromSources reads a secret from one of three mutually-exclusive
// sources -- a literal flag, an environment variable named by another flag, or
// a file path named by a third flag -- and writes the resolved value back into
// *literal so downstream code can keep reading the same pointer.
func resolveSecretFromSources(literal, envName, filePath *string, literalFlag, envFlag, fileFlag string) error {
	set := 0
	if *literal != "" {
		set++
	}
	if *envName != "" {
		set++
	}
	if *filePath != "" {
		set++
	}
	if set > 1 {
		return fmt.Errorf("%s, %s and %s are mutually exclusive; set at most one", literalFlag, envFlag, fileFlag)
	}

	if *envName != "" {
		v := os.Getenv(*envName)
		if v == "" {
			return fmt.Errorf("environment variable %q referenced by %s is empty or unset", *envName, envFlag)
		}
		*literal = v
		return nil
	}

	if *filePath != "" {
		b, err := os.ReadFile(*filePath)
		if err != nil {
			return fmt.Errorf("could not read %s %q: %w", fileFlag, *filePath, err)
		}
		*literal = strings.TrimRight(string(b), "\r\n")
		return nil
	}

	return nil
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
	c.IfDescRegStr = *interfaceDescriptionRegex
	c.InterfaceNameRegex = *interfaceNameRegex
	c.FirewallFilterNameRegex = *firewallFilterNameRegex

	f := &c.Features
	f.Accounting = *accountingEnabled
	f.Alarm = *alarmEnabled
	f.ARP = *arpEnabled
	f.BFD = *bfdEnabled
	f.BGP = *bgpEnabled
	f.Cluster = *clusterEnabled
	f.DDOSProtection = *ddosProtectionEnabled
	f.DOT1X = *dot1xEnabled
	f.Environment = *environmentEnabled
	f.EVPN = *evpnEnabled
	f.EVPNIPPrefix = *evpnIPPrefixEnabled
	f.Firewall = *firewallEnabled
	f.FPC = *fpcEnabled
	f.Interfaces = *interfacesEnabled
	f.InterfaceDiagnostic = *interfaceDiagnosticsEnabled
	f.InterfaceQueue = *interfaceQueuesEnabled
	f.IPSec = *ipsecEnabled
	f.ISIS = *isisEnabled
	f.KRT = *krtEnabled
	f.L2Circuit = *l2circuitEnabled
	f.L2Vpn = *l2vpnEnabled
	f.LACP = *lacpEnabled
	f.LDP = *ldpEnabled
	f.License = *licenseEnabled
	f.LLDP = *lldpEnabled
	f.MAC = *macEnabled
	f.MACSec = *macsecEnabled
	f.MNHA = *mnhaEnabled
	f.MPLSLSP = *mplsLSPEnabled
	f.NAT = *natEnabled
	f.NAT2 = *nat2Enabled
	f.NTP = *ntpEnabled
	f.OSPF = *ospfEnabled
	f.Poe = *poeEnabled
	f.Power = *powerEnabled
	f.Routes = *routesEnabled
	f.RoutingEngine = *routingEngineEnabled
	f.RPKI = *rpkiEnabled
	f.RPM = *rpmEnabled
	f.Satellite = *satelliteEnabled
	f.Security = *securityEnabled
	f.SecurityIKE = *securityIKEEnabled
	f.SecurityNAT = *securityNATEnabled
	f.SecurityPolicies = *securityPoliciesEnabled
	f.Storage = *storageEnabled
	f.Subscriber = *subscriberEnabled
	f.System = *systemEnabled
	f.SystemStatistics = *systemstatisticsEnabled
	f.TWAMP = *twampEnabled
	f.UFD = *ufdEnabled
	f.VirtualChassis = *virtualChassisEnabled
	f.VPWS = *vpwsEnabled
	f.VRRP = *vrrpEnabled
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

func startServer() error {
	log.Infof("Starting JunOS exporter (Version: %s)", version)
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`<html>
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

	if *webConfigFile != "" {
		log.Infof("Listening for %s on %s (web-config: %q)",
			*metricsPath, *listenAddress, *webConfigFile)
		return startListeningWithWebConfig()
	}

	log.Infof("Listening for %s on %s (TLS: %v)",
		*metricsPath, *listenAddress, *tlsEnabled)

	if *tlsEnabled {
		return http.ListenAndServeTLS(*listenAddress, *tlsCertChainPath, *tlsKeyPath, nil)
	}

	return http.ListenAndServe(*listenAddress, nil)
}

func startListeningWithWebConfig() error {
	if *tlsEnabled || *tlsCertChainPath != "" || *tlsKeyPath != "" {
		log.Warnf("-web.config.file=%q overrides legacy -tls.* flags; "+
			"TLS now comes from the YAML's tls_server_config block (or is "+
			"disabled if that block is absent)", *webConfigFile)
	}

	server := &http.Server{Addr: *listenAddress}
	flags := &web.FlagConfig{
		WebListenAddresses: &[]string{*listenAddress},
		WebSystemdSocket:   new(false),
		WebConfigFile:      webConfigFile,
	}

	return web.ListenAndServe(server, flags, slogadapter.New())
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
		http.Error(w, "POST method expected", http.StatusBadRequest)
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
		err := fmt.Errorf("logical systems not enabled but the logical system '%s' in parameters", logicalSystem)
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
