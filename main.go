package main

import (
	"fmt"
	"os"
	"passgen/subcommands"

	"github.com/spf13/cobra"
)

func main() {
	var rootVersion bool

	rootCmd := &cobra.Command{
		Use:   "passgen",
		Short: "Password generator for your daily usage",
		Long:  "Password generator for your daily usage in all your apps and websites",
		Run: func(cmd *cobra.Command, args []string) {
			if rootVersion {
				subcommands.Version.Run(cmd, args)
				return
			}

			subcommands.Generate.Run(cmd, args)
		},
	}

	rootCmd.Flags().BoolVarP(&rootVersion, "version", "V", false, "Print the version number of Passgen")

	rootCmd.AddCommand(
		subcommands.Generate,
		subcommands.Wallet,
		subcommands.Safety,
		subcommands.Version,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
