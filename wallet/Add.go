package wallet

import (
	"errors"
	"os"
	"passgen/helpers"
	"passgen/interactors"
	"unicode/utf8"

	"github.com/AlecAivazis/survey"
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
		if utf8.RuneCountInString(val.(string)) < 3 {
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

	helpers.ResultPrintf("\nYour choices are saved to wallet\n\n")
}
