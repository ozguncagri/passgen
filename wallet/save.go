package wallet

import (
	"encoding/json"
	"io/ioutil"
	"os/user"

	"github.com/ozguncagri/passgen/helpers"
)

// Save function saves marshals and saves global configuration variable to file
func Save(walletPassword string, memoryWallet *PassgenWallet) error {
	// Get current user's information
	user, err := user.Current()
	if err != nil {
		return err
	}

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
	err = ioutil.WriteFile(user.HomeDir+"/.passgen", encryptedWallet, 0777)
	if err != nil {
		return err
	}

	return nil
}
