package subcommands

import (
	"log"
	"passgen/generators"
	"passgen/helpers"
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
	// If one-time flag is passed than generate non-regeneratable passsword
	if oneTimePassword {
		pool := interactors.AskForCharPool()
		length := interactors.AskForPasswordLength()
		generatedPassword, err := generators.GenerateOneTimePassword(pool, length)
		if err != nil {
			log.Fatalln(err)
		}

		helpers.ResultPrintf("\nYour password is : %v\n\n", generatedPassword)
		return
	}

	// Otherwise generate regeneratable password
	key := interactors.AskForKeyName()
	pool := interactors.AskForCharPool()
	length := interactors.AskForPasswordLength()
	masterPassword := interactors.AskForPassword()

	generatedPassword := generators.GeneratePassword(pool, key, masterPassword, length)

	helpers.ResultPrintf("\nYour password is : %v\n\n", generatedPassword)
}
