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
* Environment (temperatures, fans and PEM power statistics)
* Routing engine statistics
* Storage (total, available and used blocks, used percentage)
* Firewall filters (counters and policers) - needs explicit rights beyond read-only
* Security policy (SRX) statistics
* Interface queue statistics
* Power (Power usage)
* License statistics (installed/used/needed)
* L2circuits (tunnel state, number of tunnels)
* LDP (number of neighbors, sessions and session states)
* VRRP (state per interface)


## Feature specific mappings
Some collected time series behave like enums - Integer values represent a certain state/meaning.

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
0: "Nonexistant"
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
`-ssh.keyfile=<file>` enables key based authentication. `-ssh.password=<password-string>` enables password based authenticaton, this can also be enabled via the config file in the form of a `password: <password-string>` entry.
Authentication order is ssh key, if none is found the cli flag is checked, the config file is checked last. If no valid auth method is specified junos_exporter exits with an error.
Specify the ssh username with the cli flag `-ssh.user`, with the `username` key under the configuration file or use the default username of `junos_exporter`.

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
    features:
      isis: true
  - host: switch\d+
    # Tell the exporter that this hostname should be used as a pattern when loading
    # device-specific configurations. This example would match against a hostname
    # like "switch123".
    host_pattern: true
    features:
      bgp: false

# Optional
# interface_description_regex: '\[([^=\]]+)(=[^\]]+)?\]'
features:
  alarm: true
  environment: true
  bgp: true
  ospf: true
  isis: false
  nat: true
  l2circuit: true
  ldp: true
  routes: true
  routing_engine: true
  firewall: false
  interfaces: true
  interface_diagnostic: true
  interface_queue: true
  storage: true
  accounting: true
  ipsec: true
  security: true
  fpc: true
  rpki: true
  rpm: false
  satellite: true
  system: true
  power: true
```

## Dynamic Interface Labels
Version 0.9.5 introduced dynamic labels retrieved from the interface descriptions. Flags are supported a well. The first part (label name) has to comply to the following rules:
* must not begin with a figure
* must only contain this charakters: A-Z,a-z,0-9,_
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

### Custom Label RegEx

To override the default behavior a `interface_description_regex` can be supplied. This parameter can be given at a global level or per device. To use per-device regexes the target devices need to be defined in the exporter config. Per-device regex cannot be used in combination with `-config.ignore-targets`.
 
#### Example
The default regex `\[([^=\]]+)(=[^\]]+)?\]` would match interface descriptions like `"Description [foo] [bar=123]"`.  
If we use `[[\s]([^=\[\]]+)(=[^,\]]+)?[,\]]` we can now match for `"Description [foo, bar=123]"` instead.  


### Grafana Dashboards

There is an example Grafana Dashboard included (grafana_dashboard.json), which has some basic variables to choose your device(s) / interface(s)

![screenshot](grafana_screenshot.png)

## Third Party Components
This software uses components of the following projects
* Prometheus Go client library (https://github.com/prometheus/client_golang)

## Contributers
for a full list of contributers have a look at https://github.com/czerwonk/junos_exporter/graphs/contributors

## License
(c) Daniel Czerwonk, 2017. Licensed under [MIT](LICENSE) license.

## Prometheus
see https://prometheus.io/

## JunOS
see https://www.juniper.net/us/en/products-services/nos/junos/
