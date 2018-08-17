package subcommands

import (
	"passgen/config"

	"github.com/spf13/cobra"
)

// WalletRemove is remotes one ore more items from wallet
func WalletRemove(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		return
	}

	for _, keyVal := range args {
		delete(config.GlobalConfig.Wallet, keyVal)
	}

	config.Save()
}
