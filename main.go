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
