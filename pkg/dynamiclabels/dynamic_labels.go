// SPDX-License-Identifier: MIT

package dynamiclabels

import (
	"regexp"
	"strings"
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

type Label struct {
	name  string
	value string
}

func (il *Label) Name() string {
	return il.name
}

func (il *Label) Value() string {
	return il.value
}

func ParseDescription(description string, ifDescReg *regexp.Regexp) Labels {
	labels := make(Labels, 0)

	if len(description) == 0 || ifDescReg == nil {
		return labels
	}

	matches := ifDescReg.FindAllStringSubmatch(description, -1)
	for _, m := range matches {
		n := strings.ToLower(m[1])

		if !nameRe.Match([]byte(n)) {
			continue
		}

		label := &Label{
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

type Labels []*Label

func (ils Labels) Keys() []string {
	ret := make([]string, 0, len(ils))
	for _, il := range ils {
		ret = append(ret, il.name)
	}

	return ret
}

func (ils Labels) Values() []string {
	ret := make([]string, 0, len(ils))
	for _, il := range ils {
		ret = append(ret, il.value)
	}

	return ret
}
