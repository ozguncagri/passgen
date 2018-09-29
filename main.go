package main

import (
	"fmt"
	"os"

	"passgen/subcommands"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "passgen",
		Short: "Password generator for your daily usage",
		Long:  "Password generator for your daily usage in all your apps and websites",
		Run:   subcommands.GenerateRunner,
	}

	rootCmd.AddCommand(
		subcommands.Generate,
		subcommands.Wallet,
		subcommands.Safety,
		subcommands.Suggestions,
		subcommands.Version,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
