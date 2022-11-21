package connector

import (
	"io"

	"golang.org/x/crypto/ssh"
)

type Device struct {
	Host string
	Auth AuthMethod
}

// AuthMethod is the method to use to authenticate agaist the device
type AuthMethod func(*ssh.ClientConfig)

// AuthByPassword uses password authentication
func AuthByPassword(username, password string) AuthMethod {
	return func(cfg *ssh.ClientConfig) {
		cfg.User = username
		cfg.Auth = append(cfg.Auth, ssh.Password(password))
	}
}

// AuthByKey uses public key authentication
func AuthByKey(username string, key io.Reader) (AuthMethod, error) {
	pk, err := loadPrivateKey(key)
	if err != nil {
		return nil, err
	}

	return func(cfg *ssh.ClientConfig) {
		cfg.User = username
		cfg.Auth = append(cfg.Auth, pk)
	}, nil
}

func (d *Device) String() string {
	return d.Host
}
