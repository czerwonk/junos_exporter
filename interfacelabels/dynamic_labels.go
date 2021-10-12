package interfacelabels

import (
	"regexp"
	"strings"
	"sync"

	"github.com/czerwonk/junos_exporter/connector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/pkg/errors"
)

var (
	nameRe *regexp.Regexp
)

func init() {
	nameRe = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
}

// NewDynamicLabels create a new instance of DynamicLabels
func NewDynamicLabels() *DynamicLabels {
	return &DynamicLabels{
		labelNames: make(map[string]int),
		labels:     make(map[interfaceKey][]*interfaceLabel),
	}
}

// DynamicLabels parses and manages dynamic labels and label values
type DynamicLabels struct {
	labelNames map[string]int
	labels     map[interfaceKey][]*interfaceLabel
	labelCount int
	mu         sync.Mutex
}

type interfaceKey struct {
	host      string
	ifaceName string
}

type interfaceLabel struct {
	name  string
	value string
}

// CollectDescriptions collects labels from descriptions
func (l *DynamicLabels) CollectDescriptions(device *connector.Device, client *rpc.Client, ifDescReg *regexp.Regexp) error {
	r := &InterfaceRPC{}
	err := client.RunCommandAndParse("show interfaces descriptions", r)
	if err != nil {
		return errors.Wrap(err, "could not retrieve interface descriptions for "+device.Host)
	}

	l.parseDescriptions(device, r.Information.Interfaces, ifDescReg)

	return nil
}

// LabelNames returns the names for all dynamic labels
func (l *DynamicLabels) LabelNames() []string {
	names := make([]string, len(l.labelNames))

	for k, v := range l.labelNames {
		names[v] = k
	}

	return names
}

// ValuesForInterface returns the values for all dynamic labels
func (l *DynamicLabels) ValuesForInterface(device *connector.Device, ifaceName string) []string {
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

func (l *DynamicLabels) parseDescriptions(device *connector.Device, ifaces []PhyInterface, ifDescReg *regexp.Regexp) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, in := range ifaces {
		labels := l.parseDescription(in, ifDescReg)

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

func (l *DynamicLabels) parseDescription(iface PhyInterface, ifDescReg *regexp.Regexp) []*interfaceLabel {
	labels := make([]*interfaceLabel, 0)

	if len(iface.Description) == 0 {
		return labels
	}

	matches := ifDescReg.FindAllStringSubmatch(iface.Description, -1)
	for _, m := range matches {
		n := strings.ToLower(m[1])

		if !nameRe.Match([]byte(n)) {
			continue
		}

		label := &interfaceLabel{
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
