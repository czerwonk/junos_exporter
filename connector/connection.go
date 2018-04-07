package connector

import (
	"bytes"
	"io/ioutil"
	"strings"

	"sync"

	"time"

	"golang.org/x/crypto/ssh"
)

const timeoutInSeconds = 5

var (
	cachedConfig *ssh.ClientConfig
	lock         = &sync.Mutex{}
)

func config(user, keyFile string) (*ssh.ClientConfig, error) {
	lock.Lock()
	defer lock.Unlock()

	if cachedConfig != nil {
		return cachedConfig, nil
	}

	pk, err := loadPublicKeyFile(keyFile)
	if err != nil {
		return nil, err
	}

	cachedConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{pk},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeoutInSeconds * time.Second,
	}

	return cachedConfig, nil
}

// NewSSSHConnection connects to device
func NewSSSHConnection(host, user, keyFile string) (*SSHConnection, error) {
	if !strings.Contains(host, ":") {
		host = host + ":22"
	}

	c := &SSHConnection{Host: host}
	err := c.Connect(user, keyFile)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// SSHConnection encapsulates the connection to the device
type SSHConnection struct {
	conn *ssh.Client
	Host string
}

// Connect connects to the device
func (c *SSHConnection) Connect(user, keyFile string) error {
	config, err := config(user, keyFile)
	if err != nil {
		return err
	}

	c.conn, err = ssh.Dial("tcp", c.Host, config)
	return err
}

// RunCommand runs a command against the device
func (c *SSHConnection) RunCommand(cmd string) ([]byte, error) {
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

// Close closes connection
func (c *SSHConnection) Close() {
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
