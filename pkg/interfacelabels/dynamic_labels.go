// SPDX-License-Identifier: MIT

package interfacelabels

import (
	"regexp"
	"strings"
	"sync"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/pkg/errors"
)

var (
	nameRe *regexp.Regexp
)

func init() {
	nameRe = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
}

func DefaultInterfaceDescRegex() *regexp.Regexp {
	return regexp.MustCompile(`\[([^=\]]+)(=[^\]]+)?\]`)
}

// NewDynamicLabelManager create a new instance of DynamicLabels
func NewDynamicLabelManager() *DynamicLabelManager {
	return &DynamicLabelManager{
		labelNames: make(map[string]int),
		labels:     make(map[interfaceKey][]*InterfaceLabel),
	}
}

// DynamicLabelManager parses and manages dynamic labels and label values
type DynamicLabelManager struct {
	labelNames map[string]int
	labels     map[interfaceKey][]*InterfaceLabel
	labelCount int
	mu         sync.Mutex
}

type interfaceKey struct {
	host      string
	ifaceName string
}

type InterfaceLabel struct {
	name  string
	value string
}

func (il *InterfaceLabel) Name() string {
	return il.name
}

func (il *InterfaceLabel) Value() string {
	return il.value
}

// CollectDescriptions collects labels from descriptions
func (l *DynamicLabelManager) CollectDescriptions(device *connector.Device, client collector.Client, ifDescReg *regexp.Regexp) error {
	r := &result{}
	err := client.RunCommandAndParse("show interfaces descriptions", r)
	if err != nil {
		return errors.Wrap(err, "could not retrieve interface descriptions for "+device.Host)
	}

	l.parseDescriptions(device, r.Information.LogicalInterfaces, ifDescReg)
	l.parseDescriptions(device, r.Information.PhysicalInterfaces, ifDescReg)

	return nil
}

// LabelNames returns the names for all dynamic labels
func (l *DynamicLabelManager) LabelNames() []string {
	names := make([]string, len(l.labelNames))

	for k, v := range l.labelNames {
		names[v] = k
	}

	return names
}

// ValuesForInterface returns the values for all dynamic labels
func (l *DynamicLabelManager) ValuesForInterface(device *connector.Device, ifaceName string) []string {
	labels := make([]string, len(l.labelNames))

	k := interfaceKey{host: device.Host, ifaceName: ifaceName}
	ifaceLabels, found := l.labels[k]
	if !found {
		return labels
	}

	for _, la := range ifaceLabels {
		labels[l.labelNames[la.name]] = la.value
	}

	return labels
}

func (l *DynamicLabelManager) parseDescriptions(device *connector.Device, ifaces []interfaceDescription, ifDescReg *regexp.Regexp) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, in := range ifaces {
		labels := ParseDescription(in.Description, ifDescReg)

		for _, la := range labels {
			if _, found := l.labelNames[la.name]; !found {
				l.labelNames[la.name] = l.labelCount
				l.labelCount++
			}

			k := interfaceKey{host: device.Host, ifaceName: in.Name}
			l.labels[k] = append(l.labels[k], la)
		}
	}
}

func ParseDescription(description string, ifDescReg *regexp.Regexp) InterfaceLabels {
	labels := make(InterfaceLabels, 0)

	if len(description) == 0 || ifDescReg == nil {
		return labels
	}

	matches := ifDescReg.FindAllStringSubmatch(description, -1)
	for _, m := range matches {
		n := strings.ToLower(m[1])

		if !nameRe.Match([]byte(n)) {
			continue
		}

		label := &InterfaceLabel{
			name: n,
		}

		val := m[2]

		if strings.HasPrefix(val, "=") {
			label.value = val[1:]
		} else {
			label.value = "1"
		}

		labels = append(labels, label)
	}

	return labels
}

type InterfaceLabels []*InterfaceLabel

func (ils InterfaceLabels) Keys() []string {
	ret := make([]string, 0, len(ils))
	for _, il := range ils {
		ret = append(ret, il.name)
	}

	return ret
}

func (ils InterfaceLabels) Values() []string {
	ret := make([]string, 0, len(ils))
	for _, il := range ils {
		ret = append(ret, il.value)
	}

	return ret
}
