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
	numberOfInterfaceLabels = 2
	prefix                  = "junos_"
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
	interfaces      []string
	interfaceLabels map[string][]string
	snmp            *gosnmp.GoSNMP
	ch              chan<- prometheus.Metric
	err             error
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

	s := &scope{interfaceLabels: make(map[string][]string), snmp: &gosnmp.GoSNMP{}, ch: ch}
	s.snmp.Port = 161
	s.snmp.Timeout = time.Duration(2) * time.Second
	s.snmp.Target = target
	s.snmp.Community = c.community
	s.snmp.Version = 1
	s.snmp.MaxOids = 255

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
	if err != nil {
		s.err = err
		return
	}
	defer s.snmp.Conn.Close()

	start := time.Now()

	err = c.fetchInterfaces(s)
	if err != nil {
		s.err = err
		return
	}

	c.fetchInterfaceLabelFromOid(".1.3.6.1.2.1.31.1.1.1.1", 0, s)
	c.fetchInterfaceLabelFromOid(".1.3.6.1.2.1.31.1.1.1.18", 1, s)

	c.fetchInterfaceMetricFromOid(".1.3.6.1.2.1.2.2.1.10", receiveBytesDesc, bitsToBytes, s)
	c.fetchInterfaceMetricFromOid(".1.3.6.1.2.1.2.2.1.16", transmitBytesDesc, bitsToBytes, s)
	c.fetchInterfaceMetricFromOid(".1.3.6.1.2.1.2.2.1.13", receiveDropsDesc, noConvert, s)
	c.fetchInterfaceMetricFromOid(".1.3.6.1.2.1.2.2.1.14", receiveErrorsDesc, noConvert, s)
	c.fetchInterfaceMetricFromOid(".1.3.6.1.2.1.2.2.1.19", transmitDropsDesc, noConvert, s)
	c.fetchInterfaceMetricFromOid(".1.3.6.1.2.1.2.2.1.20", transmitErrorsDesc, noConvert, s)

	log.Info(time.Since(start))
}

func (c *JunosCollector) fetchInterfaces(s *scope) error {
	s.interfaces = make([]string, 0)
	res, err := s.snmp.BulkWalkAll(".1.3.6.1.2.1.2.2.1.1")
	if err != nil {
		return err
	}

	for _, v := range res {
		idx := c.getId(v.Name)
		s.interfaces = append(s.interfaces, idx)
		s.interfaceLabels[idx] = make([]string, numberOfInterfaceLabels)
	}

	return nil
}

func (c *JunosCollector) fetchInterfaceLabelFromOid(oid string, index int, s *scope) {
	h := func(p gosnmp.SnmpPDU) error {
		c.handlePduAsLabel(index, p, s)
		return nil
	}

	c.fetchForInterfaces(oid, h, s)
}

func (c *JunosCollector) fetchInterfaceMetricFromOid(oid string, desc *prometheus.Desc, converter ValueConverter, s *scope) {
	h := func(p gosnmp.SnmpPDU) error {
		return c.handlePduAsMetric(desc, p, converter, s)
	}

	c.fetchForInterfaces(oid, h, s)
}

func (c *JunosCollector) fetchForInterfaces(oid string, handler func(gosnmp.SnmpPDU) error, s *scope) {
	oids := make([]string, len(s.interfaceLabels))
	i := 0
	for _, x := range s.interfaces {
		oids[i] = oid + "." + x
		i++
	}

	res, err := s.snmp.Get(oids)
	if err != nil && s.err != nil {
		s.err = err
		return
	}

	for _, v := range res.Variables {
		err := handler(v)
		if err != nil && s.err != nil {
			return
		}
	}
}

func (c *JunosCollector) handlePduAsLabel(index int, p gosnmp.SnmpPDU, s *scope) {
	id := c.getId(p.Name)

	b := p.Value.([]byte)
	s.interfaceLabels[id][index] = string(b)
}

func (c *JunosCollector) handlePduAsMetric(desc *prometheus.Desc, pdu gosnmp.SnmpPDU, converter ValueConverter, s *scope) error {
	id := c.getId(pdu.Name)
	v := pdu.Value.(uint)
	l := append(s.interfaceLabels[id], s.snmp.Target)
	m, err := prometheus.NewConstMetric(desc, prometheus.GaugeValue, converter(v), l...)

	if err != nil {
		return err
	}

	s.ch <- m

	return nil
}

func (c *JunosCollector) getId(oid string) string {
	t := strings.Split(oid, ".")
	return t[len(t)-1]
}
