// SPDX-License-Identifier: MIT

package connector

import (
	"crypto/ed25519"
	"crypto/rand"
	"net"
	"testing"

	"golang.org/x/crypto/ssh"
)

// stubSSHServer is an in-process SSH server used to drive SSHConnection through
// a real handshake. It accepts any password, replies to global requests, and
// answers "exec" requests with a fixed XML payload.
type stubSSHServer struct {
	addr     string
	listener net.Listener
	config   *ssh.ServerConfig
}

// startStubSSHServer listens on a loopback port and serves until the test ends.
func startStubSSHServer(t *testing.T) *stubSSHServer {
	t.Helper()

	_, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("could not generate host key: %v", err)
	}

	signer, err := ssh.NewSignerFromKey(priv)
	if err != nil {
		t.Fatalf("could not create signer: %v", err)
	}

	cfg := &ssh.ServerConfig{
		PasswordCallback: func(ssh.ConnMetadata, []byte) (*ssh.Permissions, error) {
			return nil, nil
		},
	}
	cfg.AddHostKey(signer)

	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("could not listen: %v", err)
	}

	s := &stubSSHServer{
		addr:     l.Addr().String(),
		listener: l,
		config:   cfg,
	}

	go s.serve()
	t.Cleanup(func() { _ = l.Close() })

	return s
}

func (s *stubSSHServer) serve() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return
		}

		go s.handleConn(conn)
	}
}

func (s *stubSSHServer) handleConn(conn net.Conn) {
	sconn, chans, reqs, err := ssh.NewServerConn(conn, s.config)
	if err != nil {
		_ = conn.Close()
		return
	}
	defer func() { _ = sconn.Close() }()

	// Replies false to keepalive@golang.org, which is enough for the client to
	// treat the connection as alive.
	go ssh.DiscardRequests(reqs)

	for newChan := range chans {
		if newChan.ChannelType() != "session" {
			_ = newChan.Reject(ssh.UnknownChannelType, "only session supported")
			continue
		}

		ch, chReqs, err := newChan.Accept()
		if err != nil {
			return
		}

		go s.handleSession(ch, chReqs)
	}
}

func (s *stubSSHServer) handleSession(ch ssh.Channel, reqs <-chan *ssh.Request) {
	defer func() { _ = ch.Close() }()

	for req := range reqs {
		if req.Type != "exec" {
			if req.WantReply {
				_ = req.Reply(false, nil)
			}
			continue
		}

		if req.WantReply {
			_ = req.Reply(true, nil)
		}

		_, _ = ch.Write([]byte("<rpc-reply></rpc-reply>"))
		_, _ = ch.SendRequest("exit-status", false, ssh.Marshal(struct{ Status uint32 }{Status: 0}))

		return
	}
}

// deviceForStub builds a Device pointing at the stub server.
func deviceForStub(s *stubSSHServer) *Device {
	return &Device{
		Host: s.addr,
		Auth: AuthByPassword("test", "test"),
	}
}
