package subcommands

import (
	"fmt"
	"log"
	"passgen/config"
	"passgen/generators"

	"github.com/spf13/cobra"
)

// WalletAdd adds password generation item to the wallet
func WalletAdd(cmd *cobra.Command, args []string) {
	key := generators.AskKeyNameForWallet()
	pool := generators.AskForCharPool()
	length := generators.AskForPasswordLength()

	config.GlobalConfig.Wallet[key] = &config.WalletItem{
		Pool:   pool,
		Length: length,
	}

	err := config.Save()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\nYour choices are saved to wallet.")
}
