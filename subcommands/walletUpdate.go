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

// WalletUpdate selects and updates wallet item
func WalletUpdate(cmd *cobra.Command, args []string) {
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

	config.GlobalConfig.Wallet[walletItemKey].Pool = generators.AskForCharPool()
	config.GlobalConfig.Wallet[walletItemKey].Length = generators.AskForPasswordLength()

	config.Save()
}
