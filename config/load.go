package config

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
)

// load loads configuration file and converts it into configuration structure
func load(config *PassgenConfig) error {
	// Get current user's information
	user, err := user.Current()
	if err != nil {
		return err
	}

	// Read and unmarshal configuration file
	if file, err := ioutil.ReadFile(user.HomeDir + "/.passgen"); err == nil {
		err = json.Unmarshal(file, &config)
		if err != nil {
			return err
		}
	}

	return nil
}
