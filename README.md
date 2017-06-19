# junos_exporter
[![Build Status](https://travis-ci.org/czerwonk/junos_exporter.svg)][travis]
[![Docker Build Statu](https://img.shields.io/docker/build/czerwonk/junos_exporter.svg)][dockerbuild]
[![Go Report Card](https://goreportcard.com/badge/github.com/czerwonk/junos_exporter)][goreportcard]

Exporter for metrics from devices running JunOS (via SNMP) https://prometheus.io/

## Remarks
this is an early version

This project is an alternative approach for collecting metrics from Juniper devices.
The set of metrics is minimal to increase performance. 
We (a few friends from the Freifunk communiy and myself) used the generic snmp_exporter before. 
Since snmp_exporter is highly generic it comes with a lot of complexity at the cost of performance. 
We wanted to have an KIS and vendor specific exporter instead. 
This approach should allow us to scrape our metrics in a very time efficient way.
For this reason this project was started.

## Install
```
go get -u github.com/czerwonk/junos_exporter
```

## Third Party Components
This software uses components of the following projects
* Prometheus Go client library (https://github.com/prometheus/client_golang)
* gosnmp (https://github.com/soniah/gosnmp)

## License
(c) Daniel Czerwonk, 2017. Licensed under [MIT](LICENSE) license.

## Prometheus
see https://prometheus.io/

[travis]: https://travis-ci.org/czerwonk/junos_exporter
[dockerbuild]: https://hub.docker.com/r/czerwonk/junos_exporter/builds
[goreportcard]: https://goreportcard.com/report/github.com/czerwonk/junos_exporter
