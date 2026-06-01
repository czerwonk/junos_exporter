// SPDX-License-Identifier: MIT

package connector

import (
	"fmt"
	"io"

	"golang.org/x/crypto/ssh"
)

func loadPrivateKey(r io.Reader, keyPassphrase string) (ssh.AuthMethod, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("could not read from reader: %w", err)
	}

	var key ssh.Signer
	if keyPassphrase == "" {
		key, err = ssh.ParsePrivateKey(b)
	} else {
		key, err = ssh.ParsePrivateKeyWithPassphrase(b, []byte(keyPassphrase))
	}

	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}

	return ssh.PublicKeys(key), nil
}
