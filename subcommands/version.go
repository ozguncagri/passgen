package subcommands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version subcommand variable
var Version = &cobra.Command{
	Use:     "version",
	Aliases: []string{"ver"},
	Short:   "Print the version number of Passgen",
	Long:    "You know every application has some version numbers. This command prints it to screen for you",
	Run:     versionRunner,
}

// version subcommand is prints application's version number to screen
func versionRunner(cmd *cobra.Command, args []string) {
	fmt.Println("Passgen : Password generator for general purpose - v3")
}
