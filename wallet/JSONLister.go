package wallet

import (
	"encoding/json"
	"fmt"
	"log"
	"passgen/config"
)

// JSONLister lists all wallet items formatted as json array
func JSONLister() {
	out, err := json.MarshalIndent(config.GlobalConfig.Wallet, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", out)
}
