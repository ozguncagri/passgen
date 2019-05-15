package storage

import (
	"os"

	"passgen/helpers"

	"github.com/olekukonko/tablewriter"
)

// List subcommand for wallet lister
func List(memoryStorage *PassgenStorage) {
	if len(memoryStorage.Storage) == 0 {
		helpers.NegativePrintf("\nThere is no item in your storage\n\n")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Application Name", "User Name", "Password"})

	for _, v := range memoryStorage.Storage {
		table.Append([]string{v.AppName, v.UserName, v.Password})
	}
	table.Render()
}
