package connector

import (
	"bytes"
	"io/ioutil"
	"strings"

	"golang.org/x/crypto/ssh"
)

func NewSshConnection(host, user, keyFile string) (*SshConnection, error) {
	if !strings.Contains(host, ":") {
		host = host + ":22"
	}

	c := &SshConnection{Host: host}
	err := c.Connect(user, keyFile)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type SshConnection struct {
	conn *ssh.Client
	Host string
}

func (c *SshConnection) Connect(user, keyFile string) error {
	pk, err := loadPublicKeyFile(keyFile)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{pk},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	c.conn, err = ssh.Dial("tcp", c.Host, config)
	return err
}

func (c *SshConnection) RunCommand(cmd string) ([]byte, error) {
	session, err := c.conn.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var b = &bytes.Buffer{}
	session.Stdout = b

	err = session.Run(cmd)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (c *SshConnection) Close() {
	if c.conn == nil {
		return
	}

	c.conn.Close()
	c.conn = nil
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
