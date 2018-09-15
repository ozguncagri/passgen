package wallet

import (
	"errors"
	"fmt"
	"os"

	"github.com/ozguncagri/passgen/helpers"
	"github.com/ozguncagri/passgen/interactors"
	"gopkg.in/AlecAivazis/survey.v1"
)

// Add adds password generation item to the wallet
func Add(memoryWallet *PassgenWallet) {
	// Ask For Wallet Key Name
	keyName := ""
	prompt := &survey.Input{
		Message: "What is key name for password :",
		Help:    "Key name should be at least 3 characters long",
	}

	err := survey.AskOne(prompt, &keyName, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) < 3 {
			return errors.New("value is too short. Min length is 3")
		}

		if _, ok := memoryWallet.Wallet[val.(string)]; ok {
			return errors.New("key is already exists in wallet")
		}

		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	pool := interactors.AskForCharPool()
	length := interactors.AskForPasswordLength()

	memoryWallet.Wallet[keyName] = Item{
		Pool:   pool,
		Length: length,
	}

	fmt.Println("\nYour choices are saved to wallet.")
}
