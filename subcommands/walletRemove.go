package subcommands

import (
	"fmt"
	"log"
	"passgen/config"

	"github.com/spf13/cobra"
)

// WalletRemove is remotes one ore more items from wallet
func WalletRemove(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		return
	}

	previousLength := len(config.GlobalConfig.Wallet)

	for _, keyVal := range args {
		delete(config.GlobalConfig.Wallet, keyVal)
	}

	if previousLength > len(config.GlobalConfig.Wallet) {
		fmt.Printf("%v key(s) are removed\n", previousLength-len(config.GlobalConfig.Wallet))
	}

	err := config.Save()
	if err != nil {
		log.Fatalln(err)
	}
}
