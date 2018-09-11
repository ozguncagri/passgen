package wallet

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
)

// Save function saves marshals and saves global configuration variable to file
func Save() error {
	// Get current user's information
	user, err := user.Current()
	if err != nil {
		return err
	}

	// Marshall current config
	configuration, err := json.Marshal(GlobalWallet)
	if err != nil {
		return err
	}

	// Write marshalled config to file
	err = ioutil.WriteFile(user.HomeDir+"/.passgen", configuration, 0777)
	if err != nil {
		return err
	}

	return nil
}
