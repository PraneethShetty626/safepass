package cmd

import (
	"fmt"
	"strings"

	"github.com/PraneethShetty626/safepasscmd/internal/utils"
	"github.com/PraneethShetty626/safepasscmd/internal/vault"

	"github.com/spf13/cobra"
)

var getCmdFlags struct {
	showPassword bool
}

var getCmd = &cobra.Command{
	Use:   "get [account] [username]",
	Short: "Retrieve an account entry",
	Long: `Retrieve details of an account or a specific username from the safepass vault.

Usage:
  safepass get [account] [username]

Examples:
  safepass get github
    Retrieves all usernames and metadata stored under the 'github' account.

  safepass get github johndoe
    Retrieves the 'johndoe' entry under the 'github' account.

Flags:
  -s, --show    Show the password in plaintext (by default, passwords are hidden)

When using this command, you'll be prompted for the master password to unlock the vault.
If the --show flag is used, passwords will be displayed in plaintext. Otherwise, they will be masked for security.`,

	Args: cobra.RangeArgs(0, 2),
	Run: func(_ *cobra.Command, args []string) {
		var account, username string

		if len(args) >= 1 {
			account = strings.TrimSpace(args[0])
		} else {
			account = strings.TrimSpace(utils.PromptLine("🔍 Enter account name: "))
			if account == "" {
				fmt.Println("❌ Error: account name cannot be empty")
				return
			}
		}

		if len(args) == 2 {
			username = strings.TrimSpace(args[1])
		}

		v := loadedVault

		pEntry, err := v.GetEntry(account)
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}

		// If username is provided, filter it
		if username != "" {
			entry, err := pEntry.GetEntry(username)
			if err != nil {
				fmt.Println("❌ Error:", err)
				return
			}
			pEntry = vault.PEntry{username: entry}
		}

		pEntry.Print(getCmdFlags.showPassword)
	},
}

func init() {
	getCmd.Flags().
		BoolVarP(&getCmdFlags.showPassword, "show", "s", false, "Show the password in plaintext")
	rootCmd.AddCommand(getCmd)
}
