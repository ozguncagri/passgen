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

	for _, v := range config.GlobalConfig.Wallet {
		allKeys = append(allKeys, v.Key)
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

	key := generators.AskKeyNameForWallet()
	pool := generators.AskForCharPool()
	length := generators.AskForPasswordLength()

	for i, v := range config.GlobalConfig.Wallet {
		if v.Key == walletItemKey {
			config.GlobalConfig.Wallet[i].Key = key
			config.GlobalConfig.Wallet[i].Pool = pool
			config.GlobalConfig.Wallet[i].Length = length
			config.Save()
			return
		}
	}
}
