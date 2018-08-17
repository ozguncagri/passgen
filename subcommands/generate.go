package subcommands

import (
	"fmt"
	"passgen/generators"

	"github.com/spf13/cobra"
)

// Generate is root command for default password generation
func Generate(cmd *cobra.Command, args []string) {
	key := generators.AskForKeyName()
	pool := generators.AskForCharPool()
	length := generators.AskForPasswordLength()
	masterPassword := generators.AskForPassword()

	generatedPassword := generators.GeneratePassword(pool, key, masterPassword, length)

	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}
