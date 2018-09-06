package subcommands

import (
	"passgen/wallet"
	"strings"

	"github.com/spf13/cobra"
)

// WalletList subcommand for wallet lister
func WalletList(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		wallet.TableLister()
		return
	}

	switch strings.ToLower(args[0]) {
	case "table":
		wallet.TableLister()
	case "json":
		wallet.JSONLister()
	case "csv":
		wallet.CSVLister()
	default:
		wallet.TableLister()
	}
}
