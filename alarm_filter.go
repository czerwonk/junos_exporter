// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"regexp"
)

// compileAlarmFilter turns the -alarms.filter value into the expression the
// alarm collector matches against. It runs at startup so an invalid expression
// stops the exporter with an error instead of reaching a scrape.
func compileAlarmFilter(expression string) (*regexp.Regexp, error) {
	if len(expression) == 0 {
		return nil, nil
	}

	re, err := regexp.Compile(expression)
	if err != nil {
		return nil, fmt.Errorf("could not compile -alarms.filter %q: %w", expression, err)
	}

	return re, nil
}
