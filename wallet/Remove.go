package wallet

import (
	"errors"
	"fmt"
	"log"
	"passgen/config"
	"passgen/helpers"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Remove is remotes one ore more items from wallet
func Remove() {
	walletItemKey := ""
	var allKeys []string

	for keys := range config.GlobalConfig.Wallet {
		allKeys = append(allKeys, keys)
	}

	prompt := &survey.Select{
		Message: "Choose wallet item for remove :",
		Options: allKeys,
	}
	err := survey.AskOne(prompt, &walletItemKey, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) == 0 {
			return errors.New("this is not valid selection")
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	previousLength := len(config.GlobalConfig.Wallet)

	delete(config.GlobalConfig.Wallet, walletItemKey)

	if previousLength > len(config.GlobalConfig.Wallet) {
		fmt.Printf("Key (%v) removed\n", walletItemKey)
	}

	err = config.Save()
	if err != nil {
		log.Fatalln(err)
	}
}
