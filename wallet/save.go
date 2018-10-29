package wallet

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"passgen/helpers"
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

	encryptedWallet, err := helpers.AES256Encrypt(walletPassword, buf.Bytes())
	if err != nil {
		return err
	}

	// Write marshalled config to file
	err = ioutil.WriteFile(walletPath, encryptedWallet, 0777)
	if err != nil {
		return err
	}

	return nil
}
