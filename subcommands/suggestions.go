package subcommands

import (
	"fmt"

	"passgen/suggestions"

	"github.com/spf13/cobra"
)

// Suggestions subcommand variable
var Suggestions = &cobra.Command{
	Use:     "suggestions",
	Aliases: []string{"advice", "info"},
	Short:   "Print suggestions for using Passgen",
	Long:    "You may need some useful suggestions for keep your account more secure with or without using Passgen",
	Args:    cobra.RangeArgs(0, 1),
	Run:     suggestionsRunner,
}

// suggestionsRunner subcommand only for giving some advices to users for secure way to use application
func suggestionsRunner(cmd *cobra.Command, args []string) {
	language := "en"

	if len(args) > 0 {
		language = args[0]
	}

	fmt.Println(suggestions.SuggestionGateway(language))
}
