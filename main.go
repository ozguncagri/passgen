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
	Run:   subcommands.Generate,
}

func init() {
	versionCommand := &cobra.Command{
		Use:     "version",
		Aliases: []string{"ver"},
		Short:   "Print the version number of Passgen",
		Long:    "You know every application has some version numbers. This command prints it to screen for you",
		Run:     subcommands.Version,
	}

	suggestionsCommand := &cobra.Command{
		Use:     "suggestions",
		Aliases: []string{"advice", "info"},
		Short:   "Print suggestions for using Passgen",
		Long:    "You may need some useful suggestions for keep your account more secure with or without using Passgen",
		Run:     subcommands.Suggestions,
	}

	generateCommand := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Password generator for your daily usage",
		Long:    "Password generator for your daily usage in all your apps and web sites",
		Run:     subcommands.Generate,
	}

	walletCommand := &cobra.Command{
		Use:   "wallet",
		Short: "Print wallet items",
		Long:  "Prints list of stored wallet items on your config",
		Run:   subcommands.Wallet,
	}

	safetyCommand := &cobra.Command{
		Use:   "safety",
		Short: "Safety tools for current passwords",
		Long:  "Sets of commands for checking current passwords safety",
	}

	safetyCommand.AddCommand(
		&cobra.Command{
			Use:   "check",
			Short: "Safety check for passwords",
			Long:  "Couple of security checks for your password's safety",
			Run:   subcommands.SafetyCheck,
		},
	)

	rootCmd.AddCommand(generateCommand, walletCommand, safetyCommand, suggestionsCommand, versionCommand)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
