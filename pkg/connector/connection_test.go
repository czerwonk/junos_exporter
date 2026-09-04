// SPDX-License-Identifier: MIT

package connector

import (
	"errors"
	"sync"
	"testing"
	"time"
)

// TestStubServerRoundTrip verifies the test harness itself: a connection can be
// established against the stub and a command runs to completion.
func TestStubServerRoundTrip(t *testing.T) {
	s := startStubSSHServer(t)

	c := NewSSHConnection(deviceForStub(s), time.Hour, time.Second)
	if err := c.Start(time.Hour); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	defer c.Stop(nil)

	b, err := c.RunCommand("show version")
	if err != nil {
		t.Fatalf("RunCommand failed: %v", err)
	}

	if got, want := string(b), "<rpc-reply></rpc-reply>"; got != want {
		t.Errorf("RunCommand output = %q, want %q", got, want)
	}

	if !c.IsConnected() {
		t.Error("IsConnected() = false after successful command, want true")
	}
}

// TestRunCommandConcurrentWithStop exercises RunCommand against a concurrent
// Stop. RunCommand reads c.sshClient without holding c.mu while Stop writes it
// under the lock, so the two accesses are unsynchronised.
func TestRunCommandConcurrentWithStop(t *testing.T) {
	s := startStubSSHServer(t)

	for round := 0; round < 20; round++ {
		c := NewSSHConnection(deviceForStub(s), time.Hour, time.Second)
		if err := c.Start(time.Hour); err != nil {
			t.Fatalf("round %d: Start failed: %v", round, err)
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for i := 0; i < 50; i++ {
				// Errors are expected once Stop lands; only the race matters.
				_, _ = c.RunCommand("show version")
			}
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			c.Stop(errors.New("concurrent stop"))
		}()

		wg.Wait()
	}
}

// TestKeepaliveConcurrentWithStop exercises the keepalive loop against a
// concurrent Stop. The loop dereferences c.tcpConn and the client returned by
// getSSHClient() without re-checking them, both of which Stop sets to nil.
func TestKeepaliveConcurrentWithStop(t *testing.T) {
	s := startStubSSHServer(t)

	// A near-zero interval keeps the keepalive timer permanently ready, so the
	// select in keepalive() actually exercises the branch where both the timer
	// and a closed done channel are runnable.
	for round := 0; round < 20; round++ {
		c := NewSSHConnection(deviceForStub(s), time.Nanosecond, time.Second)
		if err := c.Start(time.Hour); err != nil {
			t.Fatalf("round %d: Start failed: %v", round, err)
		}

		time.Sleep(time.Millisecond)
		c.Stop(errors.New("concurrent stop"))
	}

	// Give any surviving keepalive goroutine a chance to touch the nil fields.
	time.Sleep(50 * time.Millisecond)
}

// TestConnectionHandlesAfterStop pins the invariant the keepalive loop depends
// on: once a connection is stopped, the handles it would dereference are
// reported as unavailable rather than handed back as nil.
func TestConnectionHandlesAfterStop(t *testing.T) {
	s := startStubSSHServer(t)

	c := NewSSHConnection(deviceForStub(s), time.Hour, time.Second)
	if err := c.Start(time.Hour); err != nil {
		t.Fatalf("Start failed: %v", err)
	}

	tcpConn, sshClient, ok := c.connectionHandles()
	if !ok {
		t.Fatal("connectionHandles() ok = false on a live connection, want true")
	}
	if tcpConn == nil || sshClient == nil {
		t.Fatal("connectionHandles() returned a nil handle on a live connection")
	}

	c.Stop(errors.New("test"))

	tcpConn, sshClient, ok = c.connectionHandles()
	if ok {
		t.Error("connectionHandles() ok = true after Stop, want false")
	}
	if tcpConn != nil {
		t.Error("connectionHandles() returned a non-nil tcpConn after Stop")
	}
	if sshClient != nil {
		t.Error("connectionHandles() returned a non-nil sshClient after Stop")
	}
}
