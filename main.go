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
		Use:   "generate",
		Short: "Password generator for your daily usage",
		Long:  "Password generator for your daily usage in all your apps and websites",
		Run:   generalPasswordGeneratorCommand,
	}

	walletCommand := &cobra.Command{
		Use:   "wallet",
		Short: "Lists ",
		Long:  "Password generator for your daily usage in all your apps and websites",
		Run:   walletList,
	}

	walletCommand.AddCommand(
		&cobra.Command{
			Use:     "list",
			Aliases: []string{"ls"},
			Short:   "Password generator for your daily usage",
			Long:    "Password generator for your daily usage in all your apps and websites",
			Run:     walletList,
		},
		&cobra.Command{
			Use:   "add",
			Short: "Password generator for your daily usage",
			Long:  "Password generator for your daily usage in all your apps and websites",
			Run:   walletAdd,
		},
		&cobra.Command{
			Use:     "remove",
			Aliases: []string{"rm"},
			Short:   "Password generator for your daily usage",
			Long:    "Password generator for your daily usage in all your apps and websites",
			Run:     walletRemove,
		},
		&cobra.Command{
			Use:   "update",
			Short: "Password generator for your daily usage",
			Long:  "Password generator for your daily usage in all your apps and websites",
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

func walletList(cmd *cobra.Command, args []string) {

}

func walletAdd(cmd *cobra.Command, args []string) {

}

func walletRemove(cmd *cobra.Command, args []string) {

}

func walletUpdate(cmd *cobra.Command, args []string) {

}
