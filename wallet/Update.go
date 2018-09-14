package wallet

import (
	"errors"
	"fmt"
	"log"

	"github.com/ozguncagri/passgen/helpers"
	"github.com/ozguncagri/passgen/interactors"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Update selects and updates wallet item
func Update() {
	walletItemKey := ""
	var allKeys []string

	for keys := range GlobalWallet.Wallet {
		allKeys = append(allKeys, keys)
	}

	prompt := &survey.Select{
		Message: "Choose wallet item for update :",
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

	fmt.Printf("\nYou are editing : \"%v\"\n\n", walletItemKey)

	GlobalWallet.Wallet[walletItemKey] = Item{
		Pool:   interactors.AskForCharPool(),
		Length: interactors.AskForPasswordLength(),
	}

	err = Save()
	if err != nil {
		log.Fatalln(err)
	}
}
