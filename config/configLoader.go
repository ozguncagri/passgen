package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
)

// Loader loads configuration file and converts it into configuration structure
func Loader() (config PassgenConfig) {
	// Get current user's information
	user, err := user.Current()
	if err != nil {
		log.Fatalf("User info read error : %v", err)
	}

	// Read and unmarshal configuration file
	if file, err := ioutil.ReadFile(user.HomeDir + "/.passgen.json"); err == nil {
		err = json.Unmarshal(file, &config)
		if err != nil {
			log.Fatalf("Config file parse error : %v", err)
		}
	}

	return
}
