package wallet

import (
	"errors"
	"fmt"
	"log"

	"github.com/ozguncagri/passgen/helpers"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Remove is remotes one ore more items from wallet
func Remove() {
	walletItemKey := ""
	var allKeys []string

	for keys := range GlobalWallet.Wallet {
		allKeys = append(allKeys, keys)
	}

	prompt := &survey.Select{
		Message: "Choose wallet item for remove :",
		Help:    "Pick any added item for removing it from your wallet",
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

	previousLength := len(GlobalWallet.Wallet)

	delete(GlobalWallet.Wallet, walletItemKey)

	if previousLength > len(GlobalWallet.Wallet) {
		fmt.Printf("Key (%v) removed\n", walletItemKey)
	}

	err = Save()
	if err != nil {
		log.Fatalln(err)
	}
}
