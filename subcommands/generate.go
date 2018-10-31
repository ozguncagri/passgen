package subcommands

import (
	"fmt"
	"passgen/generators"
	"passgen/interactors"

	"github.com/spf13/cobra"
)

// Generate subcommand variable
var Generate = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Password generator for your daily usage",
	Long:    "Password generator for your daily usage in all your apps and web sites",
	Run:     generateRunner,
}

var oneTimePassword bool

func init() {
	Generate.Flags().BoolVarP(&oneTimePassword, "one-time", "o", false, "Generates one-time (non-regeneratable) password")
}

// generateRunner is root command for default password generation
func generateRunner(cmd *cobra.Command, args []string) {
	if oneTimePassword {
		pool := interactors.AskForCharPool()
		length := interactors.AskForPasswordLength()
		generatedPassword := generators.GenerateOneTimePassword(pool, length)

		fmt.Printf("\nYour password is : %v\n", generatedPassword)
		return
	}

	key := interactors.AskForKeyName()
	pool := interactors.AskForCharPool()
	length := interactors.AskForPasswordLength()
	masterPassword := interactors.AskForPassword()

	generatedPassword := generators.GeneratePassword(pool, key, masterPassword, length)

	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}
