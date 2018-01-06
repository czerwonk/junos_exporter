package connector

import (
	"strings"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
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

func (c *SshConnection) Close() {
	if c.session == nil {
		return
	}

	c.session.Close()
}

func openSession(conn *ssh.Client) (*ssh.Session, error) {
	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
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
