package subcommands

import (
	"fmt"
	"passgen/generators"

	"github.com/spf13/cobra"
)

func GeneralPasswordGeneratorCommand(cmd *cobra.Command, args []string) {
	appName := generators.AskForKeyName()
	password := generators.AskForPassword()
	pwdLength := generators.AskForPasswordLength()
	pool := generators.AskForCharPool()

	generatedPassword := generators.GeneratePassword(pool, appName, password, pwdLength)

	fmt.Printf("\nYour password is : %v\n", generatedPassword)
}
