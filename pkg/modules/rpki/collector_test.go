package rpki

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
)

const (
	Down = iota
	Up
	Connect
	Ex_Start
	Ex_Incr
	Ex_Full
)

func TestCollectForSession(t *testing.T) {
	var collector *rpkiCollector
	labels := []string{"target"}

	// test session
	session := session{
		IPAddress:       "217.146.23.92",
		State:           "Down",
		Flaps:           10,
		IPv4PrefixCount: 231588,
		IPv6PrefixCount: 44487,
	}

	t.Run("Down", func(t *testing.T) {
		m := &dto.Metric{}
		ch := make(chan prometheus.Metric)
		session.State = "Down"
		go collector.collectForSession(session, ch, labels)
		if err := (<-ch).Write(m); err != nil {
			fmt.Println("error write dto metric")
		}
		assert.Equal(t, int(m.Gauge.GetValue()), Down)
	})
	t.Run("Up", func(t *testing.T) {
		m := &dto.Metric{}
		ch := make(chan prometheus.Metric)
		session.State = "Up"
		go collector.collectForSession(session, ch, labels)
		if err := (<-ch).Write(m); err != nil {
			fmt.Println("error write dto metric")
		}
		assert.Equal(t, int(m.Gauge.GetValue()), Up)
	})
	t.Run("Connect", func(t *testing.T) {
		m := &dto.Metric{}
		ch := make(chan prometheus.Metric)
		session.State = "Connect"
		go collector.collectForSession(session, ch, labels)
		if err := (<-ch).Write(m); err != nil {
			fmt.Println("error write dto metric")
		}
		assert.Equal(t, int(m.Gauge.GetValue()), Connect)
	})
	t.Run("Ex_Start", func(t *testing.T) {
		m := &dto.Metric{}
		ch := make(chan prometheus.Metric)
		session.State = "Ex-Start"
		go collector.collectForSession(session, ch, labels)
		if err := (<-ch).Write(m); err != nil {
			fmt.Println("error write dto metric")
		}
		assert.Equal(t, int(m.Gauge.GetValue()), Ex_Start)
	})
	t.Run("Ex_Incr", func(t *testing.T) {
		m := &dto.Metric{}
		ch := make(chan prometheus.Metric)
		session.State = "Ex-Incr"
		go collector.collectForSession(session, ch, labels)
		if err := (<-ch).Write(m); err != nil {
			fmt.Println("error write dto metric")
		}
		assert.Equal(t, int(m.Gauge.GetValue()), Ex_Incr)
	})
	t.Run("Ex_Full", func(t *testing.T) {
		m := &dto.Metric{}
		ch := make(chan prometheus.Metric)
		session.State = "Ex-Full"
		go collector.collectForSession(session, ch, labels)
		if err := (<-ch).Write(m); err != nil {
			fmt.Println("error write dto metric")
		}
		assert.Equal(t, int(m.Gauge.GetValue()), Ex_Full)
	})
	t.Run("wrong_test", func(t *testing.T) {
		m := &dto.Metric{}
		ch := make(chan prometheus.Metric)
		session.State = "Undefined"
		go collector.collectForSession(session, ch, labels)
		if err := (<-ch).Write(m); err != nil {
			fmt.Println("error write dto metric")
		}
		assert.Equal(t, int(m.Gauge.GetValue()), Down)
	})
}
