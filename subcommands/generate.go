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
	Run:     GenerateRunner,
}

// GenerateRunner is root command for default password generation
func GenerateRunner(cmd *cobra.Command, args []string) {
	key := interactors.AskForKeyName()
	pool := interactors.AskForCharPool()
	length := interactors.AskForPasswordLength()
	masterPassword := interactors.AskForPassword()

	generatedPassword := generators.GeneratePassword(pool, key, masterPassword, length)

	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}
