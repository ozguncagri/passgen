package subcommands

import (
	"fmt"

	"passgen/suggestions"

	"github.com/spf13/cobra"
)

// Suggestions subcommand only for giving some advices to users for secure way to use application
func Suggestions(cmd *cobra.Command, args []string) {
	language := "en"

	if len(args) > 0 {
		language = args[0]
	}

	fmt.Println(suggestions.SuggestionGateway(language))
}
