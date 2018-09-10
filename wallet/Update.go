package wallet

import (
	"errors"
	"fmt"
	"log"

	"github.com/ozguncagri/passgen/config"
	"github.com/ozguncagri/passgen/generators"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Update selects and updates wallet item
func Update() {
	walletItemKey := ""
	var allKeys []string

	for keys := range config.GlobalConfig.Wallet {
		allKeys = append(allKeys, keys)
	}

	prompt := &survey.Select{
		Message: "Choose wallet item for update :",
		Options: allKeys,
	}
	err := survey.AskOne(prompt, &walletItemKey, func(val interface{}) error {
		if val.(string) == "" {
			return errors.New("this is not valid selection")
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\nYou are editing : \"%v\"\n\n", walletItemKey)

	config.GlobalConfig.Wallet[walletItemKey] = config.WalletItem{
		Pool:   generators.AskForCharPool(),
		Length: generators.AskForPasswordLength(),
	}

	err = config.Save()
	if err != nil {
		log.Fatalln(err)
	}
}
