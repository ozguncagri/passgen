package wallet

import (
	"fmt"
	"os"
	"passgen/config"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// TableLister lists all wallet items as human readable table
func TableLister() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Key Name", "Character Pool", "Length"})

	for key, v := range config.GlobalConfig.Wallet {
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

		table.Append([]string{key, strings.Join(seperatedPool, ", "), fmt.Sprint(v.Length)})
	}
	table.Render()
}
