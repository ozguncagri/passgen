package interactors

import (
	"os"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"
)

// AskForCharPool is ask for which character group are you going to use for the password you are generating
func AskForCharPool() string {
	charPools := []string{}
	prompt := &survey.MultiSelect{
		Message: "Select Character Pool Items :",
		Help:    "Selection of which character groups will be use for password generation",
		Options: []string{"Upper", "Lower", "Number", "Symbols"},
	}

	err := survey.AskOne(prompt, &charPools, survey.Required)
	if err != nil {
		os.Exit(1)
	}

	for i, v := range charPools {
		charPools[i] = string(v[0])
	}

	return strings.ToUpper(strings.Join(charPools, ""))
}
