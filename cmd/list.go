package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var listCmdFlags struct {
	showPassword bool
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all accounts and usernames stored in the vault",
	Long: `Lists all accounts and their associated usernames from the safepass vault.

Usage:
  safepass list [flags]

Flags:
  -s, --show    Show the password in plaintext (default is masked)

You'll be prompted for the master password to unlock the vault.
`,
	Args: cobra.ExactArgs(0),
	Run: func(_ *cobra.Command, _ []string) {
		// Prompt for master password
		v := loadedVault

		if len(v.Entries) == 0 {
			fmt.Println("📂 No entries found in the vault.")
			return
		}

		// Sort accounts alphabetically
		var accounts []string
		for account := range v.Entries {
			accounts = append(accounts, account)
		}
		sort.Strings(accounts)

		for _, account := range accounts {
			fmt.Printf("\n📁 Account: %s\n", account)
			pEntry := v.Entries[account]
			pEntry.Print(listCmdFlags.showPassword)
		}
	},
}

func init() {
	listCmd.Flags().
		BoolVarP(&listCmdFlags.showPassword, "show", "s", false, "Show the password in plaintext")
	rootCmd.AddCommand(listCmd)
}
