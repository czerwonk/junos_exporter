package connector

import (
	"strings"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"errors"
	"bytes"
)

func NewSshConnection(host, user, keyFile string) (*SshConnection, error) {
	if !strings.Contains(host, ":") {
		host = host+":22"
	}

	c := &SshConnection{}
	err := c.Connect(host, user, keyFile)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type SshConnection struct {
	session *ssh.Session
}

func (c *SshConnection) Connect(host, user, keyFile string) error {
	pk, err := loadPublicKeyFile(keyFile)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{pk},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return err
	}

	c.session, err = openSession(conn)
	return err
}

func (c *SshConnection) RunCommand(cmd string) ([]byte, error) {
	if c.session == nil {
		return nil, errors.New("Session must be opened first")
	}

	var b = &bytes.Buffer{}
	c.session.Stdout = b

	err := c.session.Run(cmd)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (c *SshConnection) Close() {
	if c.session == nil {
		return
	}

	c.session.Close()
	c.session = nil
}

func openSession(conn *ssh.Client) (*ssh.Session, error) {
	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func loadPublicKeyFile(file string) (ssh.AuthMethod, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(b)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(key), nil
}
