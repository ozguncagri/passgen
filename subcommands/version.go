package subcommands

import (
	"fmt"
	"passgen/config"

	"github.com/spf13/cobra"
)

// Version subcommand is prints application's version number to screen
func Version(cmd *cobra.Command, args []string) {
	fmt.Printf("Passgen : Password generator for general purpose - v%v\n", config.Version)
}
