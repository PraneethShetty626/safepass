package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"
)

// 📦 Helper: Prompt for masked password
func PromptPassword(prompt string) string {
	fmt.Print(prompt)
	passwordBytes, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return strings.TrimSpace(string(passwordBytes))
}

// 📦 Helper: Prompt for a line of input
func PromptLine(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// 📦 Helper: Get arg from CLI or prompt
func GetArgOrPrompt(args []string, index int, label string) string {
	if len(args) > index {
		return strings.TrimSpace(args[index])
	}
	return PromptLine(label + ": ")
}

func GetDefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	vaultDir := filepath.Join(home, ".safepass")
	if err := os.MkdirAll(vaultDir, 0700); err != nil {
		return "", err
	}

	return filepath.Join(vaultDir, "vault.dat"), nil
}
