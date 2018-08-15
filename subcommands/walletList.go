package subcommands

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"passgen/config"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// WalletList subcommand for wallet lister
func WalletList(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		tableLister()
		return
	}

	switch strings.ToLower(args[0]) {
	case "table":
		tableLister()
	case "json":
		jsonLister()
	case "csv":
		csvLister()
	default:
		tableLister()
	}
}

// tableLister lists all wallet items as human readable table
func tableLister() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key Name", "Character Pool", "Length"})

	for _, v := range config.GlobalConfig.Wallet {
		seperatedPool := strings.Split(v.Pool, "")

		for i := range seperatedPool {
			switch seperatedPool[i] {
			case "U":
				seperatedPool[i] = "Upper"
			case "L":
				seperatedPool[i] = "Lower"
			case "N":
				seperatedPool[i] = "Number"
			case "S":
				seperatedPool[i] = "Symbol"
			}
		}

		table.Append([]string{v.Key, strings.Join(seperatedPool, ", "), fmt.Sprint(v.Length)})
	}
	table.Render()
}

// csvLister lists all wallet items formatted as json array
func jsonLister() {
	out, err := json.MarshalIndent(config.GlobalConfig.Wallet, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", out)
}

// csvLister lists all wallet items formatted as csv
func csvLister() {
	lister := csv.NewWriter(os.Stdout)
	for _, v := range config.GlobalConfig.Wallet {
		lister.Write([]string{v.Key, v.Pool, fmt.Sprint(v.Length)})
	}
	lister.Flush()
	if err := lister.Error(); err != nil {
		log.Fatal(err)
	}
}
