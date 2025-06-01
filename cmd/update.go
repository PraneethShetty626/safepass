package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/PraneethShetty626/safepasscmd/internal/utils"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [account] [username]",
	Short: "Update the details of an account entry",
	Long: `Update the password or note for a specific username within an account.

Usage:
  safepass update <account> <username>

This command prompts you interactively to update:
- The password (optional, press Enter to keep existing)
- The note (optional, press Enter to keep existing)

Example:
  safepass update github johndoe
`,
	Args: cobra.RangeArgs(0, 2),
	Run: func(_ *cobra.Command, args []string) {
		account := utils.GetArgOrPrompt(args, 0, "Account")
		if account == "" {
			fmt.Println("❌ Error: account name cannot be empty")
			return
		}

		pEntry, err := loadedVault.GetEntry(account)
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}

		username := utils.GetArgOrPrompt(args, 1, "Username")
		if username == "" {
			fmt.Println("❌ Error: username cannot be empty")
			return
		}

		entry, err := pEntry.GetEntry(username)
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}

		fmt.Printf("Updating entry for '%s' in account '%s'\n", username, account)

		newPassword := utils.PromptLine("🔑 New password (press Enter to keep existing): ")
		newNote := utils.PromptLine("📝 New note (press Enter to keep existing): ")

		// Apply updates only if new data is provided
		if strings.TrimSpace(newPassword) != "" {
			entry.Password = newPassword
		} else {
			fmt.Println()
		}
		if strings.TrimSpace(newNote) != "" {
			entry.Note = newNote
		}

		// Update the timestamp
		entry.UpdatedAt = time.Now()

		// Save changes back
		pEntry[username] = entry

		fmt.Printf("✅ Entry for '%s' updated successfully in account '%s'.\n", username, account)
	},
	PostRun: PostSaveVault,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
