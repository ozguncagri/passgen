package subcommands

import "github.com/spf13/cobra"

// Safety subcommand variable
var Safety = &cobra.Command{
	Use:   "safety",
	Short: "Safety tools for current passwords",
	Long:  "Sets of commands for checking current passwords safety",
}

func init() {
	// Add safety check subcommand to safety command
	Safety.AddCommand(safetyCheck)
}
