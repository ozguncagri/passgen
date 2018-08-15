package subcommands

import (
	"errors"
	"fmt"
	"log"
	"passgen/config"
	"passgen/generators"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

// WalletUse is used for generating password using settings defined in wallet item
func WalletUse(cmd *cobra.Command, args []string) {
	walletItemKey := ""
	var allKeys []string

	for _, v := range config.GlobalConfig.Wallet {
		allKeys = append(allKeys, v.Key)
	}

	prompt := &survey.Select{
		Message: "Choose wallet item for generate password :",
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

	for _, v := range config.GlobalConfig.Wallet {
		if v.Key == walletItemKey {
			masterPassword := generators.AskForPassword()
			generatedPassword := generators.GeneratePassword(v.Pool, v.Key, masterPassword, v.Length)
			fmt.Printf("\nYour password is : %v\n", generatedPassword)
			return
		}
	}
}
