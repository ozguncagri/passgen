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

	var temp []config.WalletItem
	for _, v := range config.GlobalConfig.Wallet {
		isRequestFind := false
		for _, k := range args {
			if k == v.Key {
				isRequestFind = true
			}
		}

		if !isRequestFind {
			temp = append(temp, v)
		}
	}

	config.GlobalConfig.Wallet = temp
	config.Save()
}
