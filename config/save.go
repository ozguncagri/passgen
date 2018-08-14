package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
)

// Save function saves marshals and saves global configuration variable to file
func Save() {
	// Get current user's information
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	// Marshall current config
	configuration, err := json.Marshal(GlobalConfig)
	if err != nil {
		log.Fatalln(err)
	}

	// Write marshalled config to file
	err = ioutil.WriteFile(user.HomeDir+"/.passgen", configuration, 0777)
	if err != nil {
		log.Fatalln(err)
	}
}
