// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

// stoppedSignalHandler runs handleSignals and returns once it has exited, leaving
// nothing to receive from reloadCh.
func stoppedSignalHandler(t *testing.T) {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	reloadCh = make(chan chan error)
	shutdownCtx = ctx

	done := make(chan struct{})
	go func() {
		handleSignals(ctx, cancel, make(chan os.Signal, 1), make(chan os.Signal, 1))
		close(done)
	}()

	cancel()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("handleSignals did not return after its context was cancelled")
	}
}

func TestUpdateConfigurationDuringShutdown(t *testing.T) {
	stoppedSignalHandler(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/-/reload", nil)

	finished := make(chan struct{})
	go func() {
		updateConfiguration(w, req)
		close(finished)
	}()

	select {
	case <-finished:
	case <-time.After(5 * time.Second):
		t.Fatal("updateConfiguration blocked after the signal handler stopped")
	}

	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want %d", w.Code, http.StatusServiceUnavailable)
	}
}

func TestUpdateConfigurationWhenClientGivesUp(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reloadCh = make(chan chan error)
	shutdownCtx = ctx

	reqCtx, abort := context.WithCancel(context.Background())
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/-/reload", nil).WithContext(reqCtx)

	finished := make(chan struct{})
	go func() {
		updateConfiguration(w, req)
		close(finished)
	}()

	abort()

	select {
	case <-finished:
	case <-time.After(5 * time.Second):
		t.Fatal("updateConfiguration blocked after the client went away")
	}
}

// serveOneReload stands in for the signal handler, answering a single reload
// request with err.
func serveOneReload(t *testing.T, err error) {
	t.Helper()

	reloadCh = make(chan chan error)
	shutdownCtx = context.Background()

	go func() {
		rc := <-reloadCh
		rc <- err
	}()
}

func TestUpdateConfigurationReportsSuccess(t *testing.T) {
	serveOneReload(t, nil)

	w := httptest.NewRecorder()
	updateConfiguration(w, httptest.NewRequest(http.MethodPost, "/-/reload", nil))

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestUpdateConfigurationReportsReloadFailure(t *testing.T) {
	serveOneReload(t, errors.New("bad config"))

	w := httptest.NewRecorder()
	updateConfiguration(w, httptest.NewRequest(http.MethodPost, "/-/reload", nil))

	if w.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d", w.Code, http.StatusInternalServerError)
	}

	if !strings.Contains(w.Body.String(), "bad config") {
		t.Errorf("body = %q, want it to mention the reload error", w.Body.String())
	}
}
