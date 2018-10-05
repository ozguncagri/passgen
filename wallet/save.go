package wallet

import (
	"encoding/json"
	"io/ioutil"

	"passgen/helpers"
)

// Save function saves marshals and saves global configuration variable to file
func Save(walletPassword string, memoryWallet *PassgenWallet, walletPath string) error {
	// Marshall current wallet
	jsonWallet, err := json.Marshal(memoryWallet)
	if err != nil {
		return err
	}

	encryptedWallet, err := helpers.AES256Encrypt(walletPassword, jsonWallet)
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
