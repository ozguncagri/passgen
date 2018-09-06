package wallet

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"passgen/config"
)

// CSVLister lists all wallet items formatted as csv
func CSVLister() {
	lister := csv.NewWriter(os.Stdout)
	for key, v := range config.GlobalConfig.Wallet {
		lister.Write([]string{key, v.Pool, fmt.Sprint(v.Length)})
	}
	lister.Flush()
	if err := lister.Error(); err != nil {
		log.Fatal(err)
	}
}
