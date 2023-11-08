// SPDX-License-Identifier: MIT

package connector

import (
	"io"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

func loadPrivateKey(r io.Reader, keyPassphrase string) (ssh.AuthMethod, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "could not read from reader")
	}

	var key ssh.Signer
	if keyPassphrase == "" {
		key, err = ssh.ParsePrivateKey(b)
	} else {
		key, err = ssh.ParsePrivateKeyWithPassphrase(b, []byte(keyPassphrase))
	}

	if err != nil {
		return nil, errors.Wrap(err, "could not parse private key")
	}

	return ssh.PublicKeys(key), nil
}
