package main

import (
	"fmt"
	"os"
	"passgen/subcommands"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Password generator for your daily usage",
	Long:  "Password generator for your daily usage in all your apps and websites",
	Run:   subcommands.GeneralPasswordGeneratorCommand,
}

func init() {
	generateCommand := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Password generator for your daily usage",
		Long:    "Password generator for your daily usage in all your apps and web sites",
		Run:     subcommands.GeneralPasswordGeneratorCommand,
	}

	walletCommand := &cobra.Command{
		Use:   "wallet",
		Short: "Print wallet items",
		Long:  "Prints list of stored wallet items on your config",
		Run:   subcommands.WalletList,
	}

	walletCommand.AddCommand(
		&cobra.Command{
			Use:     "use",
			Aliases: []string{"gen", "generate"},
			Short:   "Generate password with wallet item",
			Long:    "Generates password with using stored wallet configuration item",
			Run:     subcommands.WalletUse,
		},
		&cobra.Command{
			Use:     "list",
			Aliases: []string{"ls"},
			Short:   "Print wallet items",
			Long:    "Prints list of stored wallet items on your config",
			Run:     subcommands.WalletList,
		},
		&cobra.Command{
			Use:   "add",
			Short: "Add new item to wallet",
			Long:  "Add new password generation configuration item to your wallet",
			Run:   subcommands.WalletAdd,
		},
		&cobra.Command{
			Use:     "remove",
			Aliases: []string{"rm"},
			Short:   "Remove an item from wallet",
			Long:    "Remove password generation configuration item from your wallet",
			Run:     subcommands.WalletRemove,
		},
		&cobra.Command{
			Use:   "update",
			Short: "Update an item on wallet",
			Long:  "Update password generation configuration item on your wallet",
			Run:   subcommands.WalletUpdate,
		},
	)

	rootCmd.AddCommand(generateCommand, walletCommand)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
