package cmd

import (
	"fmt"

	"github.com/PraneethShetty626/safepass/internal/vault"

	"github.com/PraneethShetty626/safepass/internal/utils"

	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [account] [username]",
	Short: "Add a new password entry",
	Long: `Add a new account and username entry to your encrypted safepass vault.

Examples:
  safepass add github johndoe
  safepass add gmail alice123

This command will securely prompt you for the master password, the account name, 
the username, the password, and optional notes. The entry will be encrypted and 
stored in your vault file.

Arguments:
  account   The name of the account (e.g., github, gmail)
  username  The username associated with the account

If any arguments are omitted, safepass will prompt you interactively.`,
	Args: cobra.RangeArgs(0, 2),
	Run: func(_ *cobra.Command, args []string) {

		v := loadedVault

		account := utils.GetArgOrPrompt(args, 0, "Account")
		if account == "" {
			fmt.Println("❌ Error: account name cannot be empty")
			return
		}

		username := utils.GetArgOrPrompt(args, 1, "Username")
		if username == "" {
			fmt.Println("❌ Error: username cannot be empty")
			return
		}

		password := utils.PromptPassword("🔑 Enter password: ")
		if password == "" {
			fmt.Println("❌ Error: password cannot be empty")
			return
		}

		notes := utils.PromptLine("📝 Notes (optional): ")

		entry := vault.Entry{
			Password:  password,
			UpdatedAt: time.Now(),
			Note:      notes,
		}

		if err := v.AddEntry(account, username, entry); err != nil {
			fmt.Println("❌ Error adding entry:", err)
			return
		}

		fmt.Println("✅ Entry added successfully")

	},
	PostRun: PostSaveVault,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
