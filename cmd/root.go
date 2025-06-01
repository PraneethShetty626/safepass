// safepass/cmd/root.go

/*
Copyright © 2025 praneethshetty626@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/PraneethShetty626/safepass/internal/vault"

	"github.com/PraneethShetty626/safepass/internal/utils"

	"github.com/spf13/cobra"
)

var (
	loadedVault    *vault.Vault
	masterPassword string
	vaultPath      string
)

var rootCmd = &cobra.Command{
	Use:   "safepass",
	Short: "A simple and secure password manager CLI",
	Long: `safepass is a lightweight, offline-first password manager that securely stores and retrieves your sensitive credentials.

Examples:

  safepass add <account> <username>
  safepass get <account> [username]
  safepass delete <account> [username...]

Use "safepass --help" to see available commands and options.
`,
	SilenceUsage:     true, // Suppress usage on error
	SilenceErrors:    true, // Suppress cobra's default error printing
	PersistentPreRun: PreLoadVault,
}

func Execute() {
	vPath, err := utils.GetDefaultPath()

	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	vaultPath = vPath

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "❌", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func PreLoadVault(_ *cobra.Command, _ []string) {
	masterPassword = utils.PromptPassword("🔒 Enter master password: ")

	v, err := vault.LoadVault(vaultPath, masterPassword)
	if err != nil {
		fmt.Println("❌ Error loading vault:", err)
		os.Exit(1)
	}

	loadedVault = v
}

func PostSaveVault(_ *cobra.Command, _ []string) {
	if loadedVault == nil {
		return
	}
	if err := loadedVault.SaveVault(vaultPath, masterPassword); err != nil {
		fmt.Println("❌ Error saving vault:", err)
	}
}
