package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <account> [username...]",
	Short: "Delete an account or specific usernames within an account",
	Long: `Delete an entire account or specific usernames within an account from the safepass vault.

Usage:
  safepass delete <account> [username...]

Examples:
  safepass delete github
	Deletes the entire 'github' account and all its usernames.

  safepass delete github johndoe janedoe
	Deletes only the 'johndoe' and 'janedoe' usernames from the 'github' account.

If no usernames are provided, the entire account (and all its usernames) will be deleted.
If one or more usernames are provided, only those usernames will be removed from the account.

	You'll be prompted for the master password to authorize the deletion.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		account := strings.TrimSpace(args[0])
		usernames := args[1:]

		v := loadedVault

		if len(usernames) == 0 {
			// Delete entire account
			if err := v.DeleteEntry(account); err != nil {
				fmt.Println("❌ Error:", err)
				return
			}
			fmt.Printf("✅ Account '%s' deleted successfully.\n", account)
		} else {
			// Delete specific usernames within the account
			pEntry, err := v.GetEntry(account)
			if err != nil {
				fmt.Println("❌ Error:", err)
				return
			}

			for _, username := range usernames {
				username = strings.TrimSpace(username)
				if err := pEntry.DeletePEntry(username); err != nil {
					fmt.Printf("⚠️  Error deleting '%s': %s\n", username, err)
				} else {
					fmt.Printf("✅ User '%s' deleted successfully from '%s'.\n", username, account)
				}
			}
		}

	},
	PostRun: PostSaveVault,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
