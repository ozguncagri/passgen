package subcommands

import (
	"fmt"
	"passgen/generators"

	"github.com/spf13/cobra"
)

// GeneralPasswordGeneratorCommand is default command
func GeneralPasswordGeneratorCommand(cmd *cobra.Command, args []string) {
	appName := generators.AskForKeyName()
	pool := generators.AskForCharPool()
	pwdLength := generators.AskForPasswordLength()
	password := generators.AskForPassword()

	generatedPassword := generators.GeneratePassword(pool, appName, password, pwdLength)

	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}
