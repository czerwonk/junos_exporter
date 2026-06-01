// SPDX-License-Identifier: MIT

package slogadapter

import (
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/exporter-toolkit/web"
)

// TestToolkitErrorIsForwardedAtErrorLevel drives the real exporter-toolkit
// web server with a config file that becomes invalid mid-flight, then
// makes an HTTP request to force the toolkit's webHandler.ServeHTTP into
// its error path. That path calls logger.Error(...) on the slog.Logger we
// hand it; the test asserts the call ends up at logrus's error level
// with the toolkit's structured attrs preserved as logrus fields.
func TestToolkitErrorIsForwardedAtErrorLevel(t *testing.T) {
	buf := captureLogrus(t)

	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "web.yml")

	// One bcrypt-hashed basic-auth user, no TLS -- keeps the test fast and
	// avoids generating certs. The hash is for password "s3cret"; the value
	// is irrelevant since the test never sends valid credentials.
	const goodCfg = `basic_auth_users:
  test: $2a$10$JJuSDwUuG267HEpl7mu.IuCrx8vDPF4pPvtVHMO5vYnwIOLcnrw.e
`
	if err := os.WriteFile(cfgPath, []byte(goodCfg), 0o600); err != nil {
		t.Fatal(err)
	}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer lis.Close()
	addr := lis.Addr().String()

	server := &http.Server{}
	defer server.Close()

	listenAddrs := []string{addr}
	systemdSocket := false
	flags := &web.FlagConfig{
		WebListenAddresses: &listenAddrs,
		WebSystemdSocket:   &systemdSocket,
		WebConfigFile:      &cfgPath,
	}

	serveErr := make(chan error, 1)
	go func() {
		serveErr <- web.Serve(lis, server, flags, New())
	}()

	// Short-deadline client so a stuck request can't stall the test.
	client := &http.Client{Timeout: 2 * time.Second}

	// Readiness: HTTP GET (not net.Dial) confirms the toolkit has finished
	// startup and webHandler.ServeHTTP is responding. Probing with net.Dial
	// would succeed the moment the TCP listener binds, which is *before*
	// web.Serve runs its startup-time getConfig() -- racing the corruption
	// step below could then leave us connecting to a dead listener.
	if err := waitForHTTPReady(client, "http://"+addr+"/", serveErr, 5*time.Second); err != nil {
		t.Fatalf("toolkit did not become ready: %v", err)
	}

	// Corrupt the config so the next per-request getConfig() call fails.
	if err := os.WriteFile(cfgPath, []byte("not valid yaml: [unclosed"), 0o600); err != nil {
		t.Fatal(err)
	}

	// Drop initial "Listening on" + "TLS is disabled." info lines so the
	// later assertion only sees the error line we care about.
	buf.Reset()

	resp, err := client.Get("http://" + addr + "/metrics")
	if err != nil {
		t.Fatalf("GET failed: %v", err)
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	// Toolkit logs synchronously inside ServeHTTP, but logrus may buffer at
	// the Writer boundary -- give it a brief flush window.
	time.Sleep(100 * time.Millisecond)

	got := buf.String()
	if !strings.Contains(got, "level=error") {
		t.Errorf("expected level=error in logrus output, got:\n%s", got)
	}
	if !strings.Contains(got, "Unable to parse configuration") {
		t.Errorf("expected toolkit's Error message in output, got:\n%s", got)
	}
	if !strings.Contains(got, "err=") {
		t.Errorf("expected toolkit's err= attr as a logrus field, got:\n%s", got)
	}

	// Sanity: confirm the same record at info would NOT have shown up here.
	if strings.Contains(got, "level=info") {
		t.Errorf("expected only error-level output (info lines were Reset), got:\n%s", got)
	}
}

// waitForHTTPReady polls url until an HTTP response is received (any status
// code is fine -- the point is that webHandler is installed and serving).
// Returns early with an error if web.Serve already returned, or with a
// timeout error after deadline elapses.
func waitForHTTPReady(client *http.Client, url string, serveErr <-chan error, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	var lastErr error
	for time.Now().Before(deadline) {
		select {
		case e := <-serveErr:
			if e != nil {
				return e
			}
			return nil
		default:
		}
		resp, err := client.Get(url)
		if err == nil {
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			return nil
		}
		lastErr = err
		time.Sleep(50 * time.Millisecond)
	}
	return lastErr
}
