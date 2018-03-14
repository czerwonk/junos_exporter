# junos_exporter
[![Build Status](https://travis-ci.org/czerwonk/junos_exporter.svg)](https://travis-ci.org/czerwonk/junos_exporter)
[![Docker Build Statu](https://img.shields.io/docker/build/czerwonk/junos_exporter.svg)](https://hub.docker.com/r/czerwonk/junos_exporter/builds)
[![Go Report Card](https://goreportcard.com/badge/github.com/czerwonk/junos_exporter)](https://goreportcard.com/report/github.com/czerwonk/junos_exporter)

Exporter for metrics from devices running JunOS (via SSH) https://prometheus.io/

## Remarks
This project is an alternative approach for collecting metrics from Juniper devices.
The set of metrics is minimal to increase performance. 
We (a few friends from the Freifunk communiy and myself) used the generic snmp_exporter before. 
Since snmp_exporter is highly generic it comes with a lot of complexity at the cost of performance. 
We wanted to have an KIS and vendor specific exporter instead. 
This approach should allow us to scrape our metrics in a very time efficient way.
For this reason this project was started.

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
* OSPFv3 (number of neighbors)
* Interface diagnostics (optical signals)
* ISIS (number of adjacencies, total number of routers)
* Environment (temperatures)
* Routing engine statistics

## Install
```
go get -u github.com/czerwonk/junos_exporter
```

## Usage
In this example we want to scrape 3 hosts:
* Host 1 (DNS: host1.example.com, Port: 22)
* Host 2 (DNS: host2.example.com, Port: 2233)
* Host 3 (IP: 172.16.0.1, Port: 22)

### Binary
```
./junos_exporter -ssh.targets="host1.example.com,host2.example.com:2233,172.16.0.1" -ssh.keyfile=junos_exporter
```

### Docker
```
docker run -d --restart unless-stopped -p 9326:9326 -v /opt/junos_exporter_keyfile:/ssh-keyfile -e TARGETS="host1.example.com,host2.example.com:2233,172.16.0.1" czerwonk/junos_exporter
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
