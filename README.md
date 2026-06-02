# junos_exporter
[![Go Report Card](https://goreportcard.com/badge/github.com/czerwonk/junos_exporter)](https://goreportcard.com/report/github.com/czerwonk/junos_exporter)

Exporter for metrics from devices running JunOS (via SSH) https://prometheus.io/

## Remarks
This project is an alternative approach for collecting metrics from Juniper devices.
The set of metrics is minimal to increase performance.
We (a few friends from the Freifunk community and myself) used the generic snmp_exporter before.
Since snmp_exporter is highly generic it comes with a lot of complexity at the cost of performance.
We wanted to have an KIS and vendor specific exporter instead.
This approach should allow us to scrape our metrics in a very time efficient way.
For this reason this project was started.

## Important notice for users of version < 0.10
In version 0.10 the ``config.ignore-targets`` flag was removed. The same beahior can be achieved by using an match all host pattern:
```
devices:
  - host: .*
    host_pattern: true
```

## Important notice for users of version < 0.7
In version 0.7 a typo in the prefix of all BGP related metrics was fixed. Please update your queries accordingly.

## Important notice for users of version < 0.5
In version 0.5 SNMP was replaced by SSH. This is was a breaking change (metric names were kept).
All SNMP related parameters were removed at this point.
Please have a look on the new SSH related parameters and update your service units accordingly.

## Features
The following metrics are supported by now:
* Interfaces (bytes transmitted/received, errors, drops, speed)
* Interface L1/L2 details (FEC, MAC statistics)
* L2 security (BPDU-block violations)
* Routes (per table, by protocol)
* Alarms (count)
* BGP (message count, prefix counts per peer, session state)
* OSPFv2, OSPFv3 (number of neighbors)
* Interface diagnostics (optical signals)
* ISIS (number of adjacencies, total number of routers)
* NAT (all available statistics from services nat)
* Chassis cluster HA status (SRX)
* Environment (temperatures, fans and PEM power statistics)
* EVPN (per-EVI state, neighbor route counts, detail tables for interfaces / IRBs / bridge-domains / ESIs with DF election, duplicate-MAC detection, L3 contexts)
* EVPN Type-5 / IP-prefix database (per-context per-AFI local + remote prefix counts, accepted/rejected advertisements) — separate flag (`-evpn_ip_prefix.enabled`) because the response scales with prefix count
* Routing engine statistics
* Storage (total, available and used blocks, used percentage)
* Firewall filters (counters and policers) - needs explicit rights beyond read-only
* Security policy (SRX) statistics
* Interface queue statistics
* Power (Power usage)
* License statistics (installed/used/needed)
* L2circuits (tunnel state, number of tunnels)
* LDP (number of neighbors, sessions and session states)
* LLDP (local port, local parent interface, remote port and remote system name)
* VRRP (state per interface)
* Subscribers Information (show subscribers client-type dhcp detail)
* PoE (show poe interface)
* UFD (uplink-failure-detection group state)

## Feature specific mappings
Some collected time series behave like enums - Integer values represent a certain state/meaning.

### Chassis cluster (`junos_chassis_cluster_node_status`)
```
1: primary
2: secondary
3: secondary-hold
4: disabled
5: lost
6: not-configured
7: ineligible
```

### L2circuits
```
0:EI -- encapsulation invalid
1:MM -- mtu mismatch
2:EM -- encapsulation mismatch
3:CM -- control-word mismatch
4:VM -- vlan id mismatch
5:OL -- no outgoing label
6:NC -- intf encaps not CCC/TCC
7:BK -- Backup Connection
8:CB -- rcvd cell-bundle size bad
9:LD -- local site signaled down
10:RD -- remote site signaled down
11:XX -- unknown
12:NP -- interface h/w not present
13:Dn -- down
14:VC-Dn -- Virtual circuit Down
15:Up -- operational
16:CF -- Call admission control failure
17:IB -- TDM incompatible bitrate
18:TM -- TDM misconfiguration
19:ST -- Standby Connection
20:SP -- Static Pseudowire
21:RS -- remote site standby
22:HS -- Hot-standby Connection
```

### LDP
```
0: "Nonexistent"
1: "Operational"
```

### RPKI
```
0 = "Down"
1 = "Up"
2 = "Connect"
3 = "Ex-Start"
4 = "Ex-Incr"
5 = "Ex-Full"
```

### VRRP
States map to human readable names like this:
```
1: "init"
2: "backup"
3: "master"
```

### EVPN
The EVPN collector exposes both per-EVI scalars and several state metrics. The state metrics all use the same 0/1 mapping:

```
junos_evpn_interface_status   0: Down   1: Up
junos_evpn_irb_status         0: Down   1: Up
junos_evpn_esi_resolved       0: unresolved (or remote-only)   1: resolved (local-bound)
```

`junos_evpn_esi_designated_forwarder_info` is an info-pattern gauge that is always `1` when emitted; its labels (`designated_forwarder`, `backup_forwarder`, `df_algorithm`, `local_interface`) deliberately churn on DF-election events, so it is split off from `junos_evpn_esi_resolved` to keep state-alert queries stable. Join on `(target, instance, esi)` for Grafana dashboards:

```
junos_evpn_esi_resolved
  * on(target, instance, esi)
  group_left(designated_forwarder, backup_forwarder, local_interface)
    junos_evpn_esi_designated_forwarder_info
```

The duplicate-MAC suppression total is always emitted and is the primary alert signal:
```
junos_evpn_duplicate_mac_total > 0   # forwarding loop or split-brain
```

### EVPN Type-5 IP prefix
`junos_evpn_ip_prefix_advertisement_count` uses a `status` discriminator label (`accepted`, `rejected`, …) so rejected Type-5 routes can be alerted on without separate metric families:
```
junos_evpn_ip_prefix_advertisement_count{status="rejected"} > 0
```

### License statistics
Expiry is either presented as number of days until expiry date or certain special values.
```
0 ... n = Days until expiry
     -1 = Expired
   +Inf = Permanent license
   -Inf = Invalid
```

## Install
```bash
go get -u github.com/czerwonk/junos_exporter@master
```

## Usage
In this example we want to scrape 3 hosts:
* Host 1 (DNS: host1.example.com, Port: 22)
* Host 2 (DNS: host2.example.com, Port: 2233)
* Host 3 (IP: 172.16.0.1, Port: 22)

### Binary
```bash
./junos_exporter -ssh.targets="host1.example.com,host2.example.com:2233,172.16.0.1" -ssh.keyfile=junos_exporter
```

### Docker
```bash
docker run -d --restart unless-stopped -p 9326:9326 -e SSH_KEYFILE=/ssh-keyfile -v /opt/junos_exporter_keyfile:/ssh-keyfile:ro -v /opt/junos_exporter_config.yml:/config.yml:ro czerwonk/junos_exporter
```

### Authentication
junos_exporter supports SSH authentication via key or password based authentication.
`-ssh.keyfile=<file>` enables key based authentication. `-ssh.password=<password-string>` enables password based authentication, this can also be enabled via the config file in the form of a `password: <password-string>` entry.
Authentication order is ssh key, if none is found the cli flag is checked, the config file is checked last. If no valid auth method is specified junos_exporter exits with an error.
Specify the ssh username with the cli flag `-ssh.user`, with the `username` key under the configuration file or use the default username of `junos_exporter`.

#### SSH key passphrase and password

Both the SSH key passphrase and the SSH password can be supplied from any
one of three mutually-exclusive sources -- a literal flag value, the
contents of an environment variable, or the contents of a file:

| Secret | Literal | Environment variable | File |
|---|---|---|---|
| Key passphrase | `-ssh.keyPassphrase=<string>` | `-ssh.keyPassphraseEnv=<NAME>` | `-ssh.keyPassphraseFile=<PATH>` |
| Password       | `-ssh.password=<string>`      | `-ssh.passwordEnv=<NAME>`      | `-ssh.passwordFile=<PATH>`      |

- The literal forms (`-ssh.keyPassphrase` / `-ssh.password`) are convenient
  but the value appears in `ps` output and shell history -- avoid for
  production.
- The `*Env` flags read from the named environment variable. The variable
  must exist and be non-empty at startup, otherwise the exporter exits.
- The `*File` flags read from the given path; a single trailing newline is
  trimmed. Suitable for systemd `LoadCredential=`, Docker secrets and
  Kubernetes secrets.

Within each group (passphrase, password) setting more than one source is a
configuration error and the exporter exits at startup with a clear message.
Only the global flags are affected; the per-device `key_passphrase` and
`password` fields in the YAML config are unchanged.

### Target Parameter
By default, all configured targets will be scrapped when `/metrics` is hit. As an alternative, it is possible to scrape a specific target by passing the target's hostname/IP address to the target parameter - e.g. ` http://localhost:9326/metrics?target=1.2.3.4`. The specific target must be present in the configuration file or passed in with the ssh.targets flag, you can also specify the `-config.ignore-targets` flag if you don't want to specify targets in the config or commandline, if none of this matches the request will be denied. This can be used with the below example Prometheus config:

```yaml
scrape_configs:
  - job_name: 'junos'
    static_configs:
      - targets:
        - 192.168.1.2  # Target device.
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 127.0.0.1:9326  # The junos_exporter's real hostname:port.
```

### HTTP server: TLS and basic auth

The exporter integrates [`prometheus/exporter-toolkit`](https://github.com/prometheus/exporter-toolkit),
the same web server used by `node_exporter`, `blackbox_exporter` and the rest
of the Prometheus ecosystem. Pass a web-config YAML file with
`-web.config.file=<path>` to enable TLS, HTTP basic auth, or both:

```yaml
# web.config.yml
tls_server_config:
  cert_file: /etc/junos_exporter/server.crt
  key_file:  /etc/junos_exporter/server.key

basic_auth_users:
  prom:    $2y$10$<bcrypt-hash>
  grafana: $2y$10$<another-bcrypt-hash>
```

Both blocks are optional — basic auth works over plain HTTP, and TLS works
without basic auth. The full schema (client-cert verification, cipher suites,
multiple certs, hot reload) is documented at
[exporter-toolkit/docs/web-configuration.md](https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md).

Generate a bcrypt hash with `htpasswd`:

```bash
htpasswd -B -C 10 -n -b prom 's3cret'
# prom:$2y$10$...bcrypt-hash...
```

Prometheus side:

```yaml
scrape_configs:
  - job_name: junos
    scheme: https
    basic_auth:
      username: prom
      password_file: /etc/prometheus/junos_exporter_password
    static_configs:
      - targets: [junos-exporter.example.com:9326]
```

#### Dispatch / backwards compatibility

| `-web.config.file` | `-tls.enabled` | Behaviour |
|---|---|---|
| empty | false | plain HTTP (unchanged) |
| empty | true  | TLS via legacy `-tls.cert-file` / `-tls.key-file` (unchanged) |
| set   | any   | `exporter-toolkit` handles everything; legacy `-tls.*` flags are ignored |

The legacy `-tls.enabled` / `-tls.cert-file` / `-tls.key-file` flags keep
working unchanged for existing deployments. New deployments should prefer
`-web.config.file`, which adds basic-auth support and hot-reload of cert files.

If both `-web.config.file` and any legacy `-tls.*` flag are set, the exporter
logs a warning at startup so operators see that the YAML config has taken
over and the legacy flags are being ignored.

## Config file

The exporter can be configured with a YAML based config file:

```yaml
devices:
  - host: router1
    key_file: /path/to/key
  - host: router2
    username: exporter
    password: secret
    # Optional
    # interface_description_regex: '\[([^=\]]+)(=[^\]]+)?\]'
    # interface_name_regex: '[!(d)][!(i)]*'
    # firewall_filter_name_regex: 'test-filter.*'
    features:
      isis: true
  - host: switch\d+
    # Tell the exporter that this hostname should be used as a pattern when loading
    # device-specific configurations. This example would match against a hostname
    # like "switch123".
    host_pattern: true
    features:
      bgp: false
  - host: dc-leaf\d+
    # Example: enable the EVPN collector (instance metrics, detail tables,
    # duplicate-MAC, L3 contexts) on every DC leaf. Type-5 IP-prefix routes
    # are gated separately because the response size scales with prefix count.
    host_pattern: true
    features:
      evpn: true
      evpn_ip_prefix: true

# Optional
# interface_description_regex: '\[([^=\]]+)(=[^\]]+)?\]'
# interface_name_regex: '[!(d)][!(i)]*'
# firewall_filter_name_regex: 'test-filter.*'
features:
  accounting: false
  alarm: true
  arp: false
  bfd: false
  bgp: true
  cluster: false
  ddos_protection: false
  dot1x: false
  environment: true
  evpn: false
  evpn_ip_prefix: false
  firewall: true
  fpc: false
  interface_diagnostic: true
  interface_queue: true
  interfaces: true
  ipsec: false
  isis: true
  krt: false
  l2circuit: false
  l2vpn: false
  lacp: false
  ldp: true
  license: false
  lldp: false
  mac: false
  macsec: true
  mpls_lsp: false
  nat: false
  nat2: false
  ntp: false
  ospf: true
  poe: false
  power: false
  routes: true
  routing_engine: true
  rpki: false
  rpm: false
  satellite: false
  security: false
  security_ike: false
  security_policies: false
  storage: false
  subscriber: false
  system: false
  system_statistics: true
  twamp: false
  ufd: false
  vpws: false
  vrrp: false
```

## Dynamic Interface Labels
Version 0.9.5 introduced dynamic labels retrieved from the interface descriptions. Version 0.12.4 added support for dynamic labels on BGP metrics. Flags are supported a well. The first part (label name) has to comply to the following rules:
* must not begin with a figure
* must only contain this characters: A-Z,a-z,0-9,_
* is treated lower case
* must no conflict with label names used in junos_exporter

Values can contain arbitrary characters.

The complete feature can be disabled by setting ``-dynamic-interface-labels`` to false.

### Examples
Tags:
```
Description: XYZ [prod]
Label name: prod
Label value: 1
```

Label value pairs:
```
Description: XYZ [peer=202739]
Label name: peer
Label value: 202739
```

### Enriching other collectors via PromQL join
Dynamic labels are attached only to metrics whose source RPC carries the description text (interfaces, interfacediagnostics, interfacequeue, bgp). Other collectors with an `interface` label — for example EVPN's `junos_evpn_interface_status` or `junos_evpn_esi_designated_forwarder_info` — can pick up the same labels at query time via a vector join on `(target, interface)`:
```
junos_evpn_interface_status
  * on(target, interface) group_left(prod, peer, customer)
    junos_interfaces_up
```
This requires the `interfaces` collector to be enabled (the default).

### Custom Label RegEx

To override the default behavior a `interface_description_regex` can be supplied. This parameter can be given at a global level or per device. To use per-device regexes the target devices need to be defined in the exporter config. Per-device regex cannot be used in combination with `-config.ignore-targets`.

#### Example
The default regex `\[([^=\]]+)(=[^\]]+)?\]` would match interface descriptions like `"Description [foo] [bar=123]"`.
If we use `[[\s]([^=\[\]]+)(=[^,\]]+)?[,\]]` we can now match for `"Description [foo, bar=123]"` instead.


### Configuring Interfaces Collector Command Argument

By default, the interfaces collector executes the command `show interfaces extensive` to retrieve detailed interface statistics.
If you want to query only specific interfaces or apply a wildcard filter (for example, to reduce scrape times or target specific ports), you can configure a custom argument using the `interface_name_regex` option. One example is for use with subscriber management, where `'"[!(d)][!(i)]*"'` will avoid scraping all the demux interfaces.

This argument can be supplied via:
- **CLI Flag**: `-interfaces.name-regex="\"[!(d)][!(i)]*\""`
- **Config File (Global)**: `interface_name_regex: '"[!(d)][!(i)]*"'`
- **Config File (Per-Device)**: Under a specific device configuration:
  ```yaml
  devices:
    - host: router1
      interface_name_regex: 'ge-*'
  ```

If provided, the exporter will execute `show interfaces <argument> extensive` instead of the default `show interfaces extensive`.


### Configuring Firewall Filter Name Regex

By default, the firewall collector executes the command `show firewall filter regex .*` to retrieve statistics for all firewall filters.
If you want to limit the collector to query only specific filters matching a regular expression, you can configure a custom regex using the `firewall_filter_name_regex` option.

This regex can be supplied via:
- **CLI Flag**: `-firewall.filter-name-regex="test-filter.*"`
- **Config File (Global)**: `firewall_filter_name_regex: 'test-filter.*'`
- **Config File (Per-Device)**: Under a specific device configuration:
  ```yaml
  devices:
    - host: router1
      firewall_filter_name_regex: 'router1-filter.*'
  ```

If provided, the exporter will execute `show firewall filter regex <regex>` with your custom pattern instead of `.*`.


### Grafana Dashboards
There are example Grafana dashboards included in [example/dashboards](example/dashboards).

## Third Party Components
This software uses components of the following projects
* Prometheus Go client library (https://github.com/prometheus/client_golang)

## Contributors
for a full list of contributors have a look at https://github.com/czerwonk/junos_exporter/graphs/contributors

## License
(c) Daniel Czerwonk, 2017. Licensed under [MIT](LICENSE) license.

## Prometheus
see https://prometheus.io/

## JunOS
see https://www.juniper.net/us/en/products-services/nos/junos/
