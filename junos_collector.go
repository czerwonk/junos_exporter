package main

import (
	"strings"

	"sync"

	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/soniah/gosnmp"
)

type ValueConverter func(uint) float64

const (
	numberOfLabels = 2
	prefix         = "junos_"
)

var (
	upDesc             *prometheus.Desc
	receiveBytesDesc   *prometheus.Desc
	receiveErrorsDesc  *prometheus.Desc
	receiveDropsDesc   *prometheus.Desc
	transmitBytesDesc  *prometheus.Desc
	transmitErrorsDesc *prometheus.Desc
	transmitDropsDesc  *prometheus.Desc
)

func init() {
	upDesc = prometheus.NewDesc(prefix+"up", "Scrape of target was successful", []string{"target"}, nil)

	l := []string{"name", "description", "target"}
	receiveBytesDesc = prometheus.NewDesc(prefix+"interface_receive_bytes", "Received data in bytes", l, nil)
	receiveErrorsDesc = prometheus.NewDesc(prefix+"interface_receive_errors", "Number of errors caused by incoming packets", l, nil)
	receiveDropsDesc = prometheus.NewDesc(prefix+"interface_receive_drops", "Number of dropped incoming packets", l, nil)
	transmitBytesDesc = prometheus.NewDesc(prefix+"interface_transmit_bytes", "Transmitted data in bytes", l, nil)
	transmitErrorsDesc = prometheus.NewDesc(prefix+"interface_transmit_errors", "Number of errors caused by outgoing packets", l, nil)
	transmitDropsDesc = prometheus.NewDesc(prefix+"interface_transmit_drops", "Number of dropped outgoing packets", l, nil)
}

type JunosCollector struct {
	targets   []string
	community string
}

type scope struct {
	labelValues map[string][]string
	snmp        *gosnmp.GoSNMP
	ch          chan<- prometheus.Metric
	err         error
}

func NewJunosCollector(targets []string, community string) *JunosCollector {
	return &JunosCollector{targets: targets, community: community}
}

func (c *JunosCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- receiveBytesDesc
	ch <- receiveErrorsDesc
	ch <- receiveDropsDesc
	ch <- transmitBytesDesc
	ch <- transmitDropsDesc
	ch <- transmitErrorsDesc
}

func (c *JunosCollector) Collect(ch chan<- prometheus.Metric) {
	wg := &sync.WaitGroup{}
	wg.Add(len(c.targets))

	for _, t := range c.targets {
		go c.collectForTarget(t, ch, wg)
	}

	wg.Wait()
}

func (c *JunosCollector) collectForTarget(target string, ch chan<- prometheus.Metric, wg *sync.WaitGroup) {
	defer wg.Done()

	s := &scope{labelValues: make(map[string][]string), snmp: &gosnmp.GoSNMP{}, ch: ch}
	s.snmp.Port = 161
	s.snmp.Timeout = time.Duration(2) * time.Second
	s.snmp.Target = target
	s.snmp.Community = c.community
	s.snmp.Version = 1

	c.collectMetrics(s)
	if s.err != nil {
		log.Error(s.err)

		ch <- c.upMetric(0, s)
		return
	}

	ch <- c.upMetric(1, s)
}

func (c *JunosCollector) upMetric(value float64, s *scope) prometheus.Metric {
	m, _ := prometheus.NewConstMetric(upDesc, prometheus.GaugeValue, value, s.snmp.Target)
	return m
}

func (c *JunosCollector) collectMetrics(s *scope) {
	err := s.snmp.Connect()

	if err != nil && s.err == nil {
		s.err = err
		return
	}

	defer s.snmp.Conn.Close()

	c.fetchLabelFromOid(".1.3.6.1.2.1.31.1.1.1.1", 0, s)
	c.fetchLabelFromOid(".1.3.6.1.2.1.31.1.1.1.18", 1, s)

	c.fetchMetricFromOid(".1.3.6.1.2.1.2.2.1.10", receiveBytesDesc, bitsToBytes, s)
	c.fetchMetricFromOid(".1.3.6.1.2.1.2.2.1.16", transmitBytesDesc, bitsToBytes, s)
	c.fetchMetricFromOid(".1.3.6.1.2.1.2.2.1.13", receiveDropsDesc, noConvert, s)
	c.fetchMetricFromOid(".1.3.6.1.2.1.2.2.1.14", receiveErrorsDesc, noConvert, s)
	c.fetchMetricFromOid(".1.3.6.1.2.1.2.2.1.19", transmitDropsDesc, noConvert, s)
	c.fetchMetricFromOid(".1.3.6.1.2.1.2.2.1.20", transmitErrorsDesc, noConvert, s)
}

func (c *JunosCollector) fetchLabelFromOid(oid string, index int, s *scope) {
	err := s.snmp.Walk(oid, func(pdu gosnmp.SnmpPDU) error {
		c.handlePduAsLabel(index, pdu, s)
		return nil
	})

	if err != nil && s.err == nil {
		s.err = err
	}
}

func (c *JunosCollector) handlePduAsLabel(index int, p gosnmp.SnmpPDU, s *scope) {
	id := c.getId(p.Name)

	l, found := s.labelValues[id]
	if !found {
		l = make([]string, numberOfLabels)
		s.labelValues[id] = l
	}

	b := p.Value.([]byte)
	l[index] = string(b)
}

func (c *JunosCollector) getId(oid string) string {
	t := strings.Split(oid, ".")
	return t[len(t)-1]
}

func (c *JunosCollector) fetchMetricFromOid(oid string, desc *prometheus.Desc, converter ValueConverter, s *scope) {
	err := s.snmp.Walk(oid, func(pdu gosnmp.SnmpPDU) error {
		return c.handlePduAsMetric(desc, pdu, converter, s)
	})

	if err != nil && s.err == nil {
		s.err = err
	}
}

func (c *JunosCollector) handlePduAsMetric(desc *prometheus.Desc, pdu gosnmp.SnmpPDU, converter ValueConverter, s *scope) error {
	id := c.getId(pdu.Name)
	v := pdu.Value.(uint)
	l := append(s.labelValues[id], s.snmp.Target)
	m, err := prometheus.NewConstMetric(desc, prometheus.GaugeValue, converter(v), l...)

	if err != nil {
		return err
	}

	s.ch <- m

	return nil
}
