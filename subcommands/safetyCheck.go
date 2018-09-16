package subcommands

import (
	"errors"
	"os"
	"strings"

	"passgen/helpers"
	"passgen/safety"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

// SafetyCheck runs couple of security checks for your password's safety
func SafetyCheck(cmd *cobra.Command, args []string) {
	password := ""
	prompt := &survey.Password{
		Message: "What is your password :",
		Help:    "Write any of your current password for security check",
	}

	err := survey.AskOne(prompt, &password, func(val interface{}) error {
		if helpers.ProperCharacterCounter(val.(string)) < 1 {
			return errors.New("value is too short. Min length is 1")
		}
		return nil
	})
	if err != nil {
		os.Exit(1)
	}

	safety.Check(strings.TrimSpace(password))
}
