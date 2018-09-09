package subcommands

import (
	"errors"
	"log"
	"passgen/helpers"
	"passgen/wallet"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

// Wallet adds password generation item to the wallet
func Wallet(cmd *cobra.Command, args []string) {
	walletItemKey := ""
	allKeys := []string{"Generate", "Add", "List", "Update", "Remove"}

	prompt := &survey.Select{
		Message: "What do you want to do on your wallet :",
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

	switch walletItemKey {
	case "Generate":
		wallet.Use()
	case "Add":
		wallet.Add()
	case "List":
		wallet.List()
	case "Update":
		wallet.Update()
	case "Remove":
		wallet.Remove()
	}
}
