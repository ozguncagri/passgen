package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Password generator for your daily usage",
	Long:  "Password generator for your daily usage in all your apps and websites",
	Run:   generalPasswordGeneratorCommand,
}

func init() {
	generateCommand := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Password generator for your daily usage",
		Long:    "Password generator for your daily usage in all your apps and web sites",
		Run:     generalPasswordGeneratorCommand,
	}

	walletCommand := &cobra.Command{
		Use:   "wallet",
		Short: "Print wallet items",
		Long:  "Prints list of stored wallet items on your config",
		Run:   walletList,
	}

	walletCommand.AddCommand(
		&cobra.Command{
			Use:     "use",
			Aliases: []string{"gen", "generate"},
			Short:   "Generate password with wallet item",
			Long:    "Generates password with using stored wallet configuration item",
			Run:     walletUse,
		},
		&cobra.Command{
			Use:     "list",
			Aliases: []string{"ls"},
			Short:   "Print wallet items",
			Long:    "Prints list of stored wallet items on your config",
			Run:     walletList,
		},
		&cobra.Command{
			Use:   "add",
			Short: "Add new item to wallet",
			Long:  "Add new password generation configuration item to your wallet",
			Run:   walletAdd,
		},
		&cobra.Command{
			Use:     "remove",
			Aliases: []string{"rm"},
			Short:   "Remove an item from wallet",
			Long:    "Remove password generation configuration item from your wallet",
			Run:     walletRemove,
		},
		&cobra.Command{
			Use:   "update",
			Short: "Update an item on wallet",
			Long:  "Update password generation configuration item on your wallet",
			Run:   walletUpdate,
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

func generalPasswordGeneratorCommand(cmd *cobra.Command, args []string) {
	appName := askForKeyName()
	password := askForPassword()
	pwdLength := askForPasswordLength()
	pool := askForCharPool()

	generatedPassword := generatePassword(pool, appName, password, pwdLength)

	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}

func walletUse(cmd *cobra.Command, args []string) {

}

func walletList(cmd *cobra.Command, args []string) {

}

func walletAdd(cmd *cobra.Command, args []string) {

}

func walletRemove(cmd *cobra.Command, args []string) {

}

func walletUpdate(cmd *cobra.Command, args []string) {

}
