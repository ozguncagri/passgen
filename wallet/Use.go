package wallet

import (
	"errors"
	"fmt"
	"log"

	"github.com/ozguncagri/passgen/generators"

	"github.com/ozguncagri/passgen/interactors"

	"github.com/ozguncagri/passgen/helpers"

	"gopkg.in/AlecAivazis/survey.v1"
)

// Use is used for generating password using settings defined in wallet item
func Use() {
	walletItemKey := ""
	var allKeys []string

	for keys := range GlobalWallet.Wallet {
		allKeys = append(allKeys, keys)
	}

	prompt := &survey.Select{
		Message: "Choose wallet item for generate password :",
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

	masterPassword := interactors.AskForPassword()
	generatedPassword := generators.GeneratePassword(
		GlobalWallet.Wallet[walletItemKey].Pool,
		walletItemKey,
		masterPassword,
		GlobalWallet.Wallet[walletItemKey].Length,
	)
	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}
