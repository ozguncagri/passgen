package suggestions

import (
	"fmt"
	"log"
	"os/user"
	"strings"
)

var allSuggestions = make(map[string]string)

// SuggestionGateway brings suggestion text with selected language
func SuggestionGateway(languageAbbreviation string) string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	name := strings.Trim(strings.Split(user.Name, " ")[0], " ")
	if len(name) == 0 {
		name = user.Username
	}

	if suggestion, ok := allSuggestions[languageAbbreviation]; ok {
		return fmt.Sprintf(suggestion, name)
	}

	return fmt.Sprintf(allSuggestions["en"], name)
}
