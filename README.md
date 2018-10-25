# junos_exporter
[![Build Status](https://travis-ci.org/czerwonk/junos_exporter.svg)](https://travis-ci.org/czerwonk/junos_exporter)
[![Docker Build Statu](https://img.shields.io/docker/build/czerwonk/junos_exporter.svg)](https://hub.docker.com/r/czerwonk/junos_exporter/builds)
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

## Important notice for users of version < 0.7
In version 0.7 a typo in the prefix of all BGP related metrics was fixed. Please update your queries accordingly. 

## Important notice for users of version < 0.5
In version 0.5 SNMP was replaced by SSH. This is was a breaking change (metric names were kept). 
All SNMP related parameters were removed at this point. 
Please have a look on the new SSH related parameters and update your service units accordingly.

## Features
The following metrics are supported by now:
* Interfaces (bytes transmitted/received, errors, drops)
* Routes (per table, by protocol)
* Alarms (count)
* BGP (message count, prefix counts per peer, session state)
* OSPFv2, OSPFv3 (number of neighbors)
* Interface diagnostics (optical signals)
* ISIS (number of adjacencies, total number of routers)
* Environment (temperatures)
* Routing engine statistics

## Install
```bash
go get -u github.com/czerwonk/junos_exporter
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
docker run -d --restart unless-stopped -p 9326:9326 -v /opt/junos_exporter_keyfile:/ssh-keyfile:ro -v /opt/junos_exporter_config.yml:/config.yml:ro czerwonk/junos_exporter
```

### Target Parameter
By default, all configured targets will be scrapped when `/metrics` is hit. As an alternative, it is possible to scrape a specific target by passing the target's hostname/IP address to the target parameter - e.g. ` http://localhost:9326/metrics?target=1.2.3.4`. The specific target must be present in the configuration file or passed in with the ssh.targets flag, otherwise, the request will be denied. This can be used with the below example Prometheus config:

```
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
targets:
  - router1
  - router2

features:
  bgp: true
  ospf: false
  isis: false
  environment: true
  routes: true
  routing_engine: true
  interface_diagnostic: true
```

## Third Party Components
This software uses components of the following projects
* Prometheus Go client library (https://github.com/prometheus/client_golang)

## Contributers
Martin (https://github.com/l3akage)

## License
(c) Daniel Czerwonk, 2017. Licensed under [MIT](LICENSE) license.

## Prometheus
see https://prometheus.io/

## JunOS
see https://www.juniper.net/us/en/products-services/nos/junos/
