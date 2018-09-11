package wallet

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
)

// Load loads configuration file and converts it into configuration structure
func Load() error {
	// Get current user's information
	user, err := user.Current()
	if err != nil {
		return err
	}

	// Read and unmarshal configuration file
	if file, err := ioutil.ReadFile(user.HomeDir + "/.passgen"); err == nil {

		err = json.Unmarshal(file, &GlobalWallet)
		if err != nil {
			return err
		}
	}

	return nil
}
