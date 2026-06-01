// SPDX-License-Identifier: MIT

package slogadapter

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

// captureLogrus redirects the package-level logrus logger into a buffer at
// Debug level for the duration of the test. It restores the previous
// settings on cleanup.
func captureLogrus(t *testing.T) *bytes.Buffer {
	t.Helper()
	prevOut := log.StandardLogger().Out
	prevLvl := log.StandardLogger().Level
	t.Cleanup(func() {
		log.StandardLogger().Out = prevOut
		log.StandardLogger().Level = prevLvl
	})
	var buf bytes.Buffer
	log.StandardLogger().Out = &buf
	log.StandardLogger().SetLevel(log.DebugLevel)
	return &buf
}

func TestLevelMapping(t *testing.T) {
	cases := []struct {
		slogLevel  slog.Level
		wantOutput string
	}{
		{slog.LevelDebug, "level=debug"},
		{slog.LevelInfo, "level=info"},
		{slog.LevelWarn, "level=warning"},
		{slog.LevelError, "level=error"},
	}

	for _, tc := range cases {
		t.Run(tc.slogLevel.String(), func(t *testing.T) {
			buf := captureLogrus(t)

			slog.New(&Handler{}).Log(context.Background(), tc.slogLevel, "hello", "k", "v")

			got := buf.String()
			if !strings.Contains(got, tc.wantOutput) {
				t.Errorf("level %s: expected %q in output, got: %s",
					tc.slogLevel, tc.wantOutput, got)
			}
			if !strings.Contains(got, "k=v") {
				t.Errorf("level %s: expected attr k=v in output, got: %s",
					tc.slogLevel, got)
			}
			if !strings.Contains(got, `msg=hello`) && !strings.Contains(got, `msg="hello"`) {
				t.Errorf("level %s: expected msg=hello in output, got: %s",
					tc.slogLevel, got)
			}
		})
	}
}

func TestEnabledRespectsLogrusLevel(t *testing.T) {
	prevLvl := log.StandardLogger().Level
	t.Cleanup(func() { log.StandardLogger().SetLevel(prevLvl) })

	h := &Handler{}

	log.StandardLogger().SetLevel(log.WarnLevel)
	if h.Enabled(context.Background(), slog.LevelInfo) {
		t.Error("Info should be disabled when logrus is at Warn")
	}
	if !h.Enabled(context.Background(), slog.LevelError) {
		t.Error("Error should be enabled when logrus is at Warn")
	}

	log.StandardLogger().SetLevel(log.DebugLevel)
	if !h.Enabled(context.Background(), slog.LevelDebug) {
		t.Error("Debug should be enabled when logrus is at Debug")
	}
}

func TestAttrsAreForwardedAsFields(t *testing.T) {
	buf := captureLogrus(t)

	slog.New(&Handler{}).Info("listening",
		slog.String("address", ":9326"),
		slog.Bool("tls", true),
		slog.Int("port", 9326),
	)

	got := buf.String()
	wantSubstrings := []string{
		`msg=listening`,
		`address=":9326"`,
		`tls=true`,
		`port=9326`,
	}
	for _, want := range wantSubstrings {
		if !strings.Contains(got, want) {
			t.Errorf("expected %q in output, got: %s", want, got)
		}
	}
}

func TestWithAttrsPersistsAcrossCalls(t *testing.T) {
	buf := captureLogrus(t)

	logger := slog.New(&Handler{}).With(slog.String("component", "web"))
	logger.Info("first")
	logger.Info("second", slog.String("path", "/metrics"))

	got := buf.String()
	if c := strings.Count(got, `component=web`); c != 2 {
		t.Errorf("expected component=web on both lines (got %d): %s", c, got)
	}
	if !strings.Contains(got, `path=/metrics`) {
		t.Errorf("expected path=/metrics in second line, got: %s", got)
	}
}

func TestWithGroupPrefixesKeys(t *testing.T) {
	buf := captureLogrus(t)

	slog.New(&Handler{}).WithGroup("server").Info("up",
		slog.String("address", ":9326"),
		slog.Bool("tls", true),
	)

	got := buf.String()
	wantSubstrings := []string{
		`server.address`,
		`server.tls=true`,
	}
	for _, want := range wantSubstrings {
		if !strings.Contains(got, want) {
			t.Errorf("expected %q in output, got: %s", want, got)
		}
	}
}

func TestNestedGroupAttrPrefixesKeys(t *testing.T) {
	buf := captureLogrus(t)

	slog.New(&Handler{}).Info("up",
		slog.Group("server",
			slog.String("address", ":9326"),
			slog.Group("tls",
				slog.Bool("enabled", true),
				slog.String("cert", "server.crt"),
			),
		),
	)

	got := buf.String()
	wantSubstrings := []string{
		`server.address`,
		`server.tls.enabled=true`,
		`server.tls.cert=server.crt`,
	}
	for _, want := range wantSubstrings {
		if !strings.Contains(got, want) {
			t.Errorf("expected %q in output, got: %s", want, got)
		}
	}
}

type lazyValue struct{ s string }

func (lv lazyValue) LogValue() slog.Value { return slog.StringValue(lv.s) }

func TestLogValuerResolved(t *testing.T) {
	buf := captureLogrus(t)

	slog.New(&Handler{}).Info("hello", slog.Any("v", lazyValue{s: "resolved"}))

	got := buf.String()
	if !strings.Contains(got, `v=resolved`) {
		t.Errorf("expected LogValuer to resolve to 'resolved', got: %s", got)
	}
}

func TestEmptyAttrDropped(t *testing.T) {
	buf := captureLogrus(t)

	// An Attr with an empty Key and a non-group value should be dropped per
	// slog's documented rules.
	slog.New(&Handler{}).Info("hello", slog.Attr{Key: "", Value: slog.StringValue("nope")})

	got := buf.String()
	if strings.Contains(got, "nope") {
		t.Errorf("expected empty-key attr to be dropped, got: %s", got)
	}
}
