package subcommands

import (
	"fmt"

	"github.com/ozguncagri/passgen/generators"
	"github.com/ozguncagri/passgen/interactors"

	"github.com/spf13/cobra"
)

// Generate is root command for default password generation
func Generate(cmd *cobra.Command, args []string) {
	key := interactors.AskForKeyName()
	pool := interactors.AskForCharPool()
	length := interactors.AskForPasswordLength()
	masterPassword := interactors.AskForPassword()

	generatedPassword := generators.GeneratePassword(pool, key, masterPassword, length)

	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}
