package wallet

import (
	"bytes"
	"encoding/gob"
	"os"
	"passgen/helpers"
	"path/filepath"
)

// Save function saves marshals and saves global configuration variable to file
func Save(walletPassword string, memoryWallet *PassgenWallet, walletPath string) error {
	// Define encode buffer
	var buf bytes.Buffer

	// Define encoder
	encoder := gob.NewEncoder(&buf)

	// Serialize wallet object
	err := encoder.Encode(memoryWallet)
	if err != nil {
		return err
	}

	// Encrypt encoded wallet
	encryptedWallet, err := helpers.Encrypt(buf.Bytes(), walletPassword)
	if err != nil {
		return err
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(walletPath), 0o600); err != nil {
		return err
	}

	// Write encrypted wallet to file
	err = os.WriteFile(walletPath, encryptedWallet, 0o600)
	if err != nil {
		return err
	}

	return nil
}
