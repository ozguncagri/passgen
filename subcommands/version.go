package subcommands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version subcommand is prints application's version number to screen
func Version(cmd *cobra.Command, args []string) {
	fmt.Println("Passgen : Password generator for general purpose - v1")
}
