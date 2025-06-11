# SafePass 🔐

A lightweight, offline-first password manager that securely stores and retrieves your sensitive credentials via command line interface.

## Features

- **Offline-First**: All data stored locally, no internet connection required
- **Secure Encryption**: Passwords encrypted with master password protection
- **Simple CLI Interface**: Fast and intuitive command-line operations
- **Account-Based Organization**: Group credentials by accounts (e.g., github, gmail)
- **Multiple Users per Account**: Store multiple usernames for the same service
- **Optional Notes**: Add notes to password entries for additional context
- **Password Masking**: Passwords hidden by default, show with `--show` flag
- **Cross-Platform**: Built with Go for Windows, macOS, and Linux compatibility

## Installation

### Prerequisites

- Go 1.19 or higher (for building from source)

### Install from Source

```bash
git clone https://github.com/PraneethShetty626/safepass.git
cd safepass
go build -o safepass
```

### Add to PATH (Optional)

```bash
# Linux/macOS
sudo mv safepass /usr/local/bin/

# Or add to your PATH
export PATH=$PATH:/path/to/safepass
```

## Quick Start

### Add Your First Password

```bash
safepass add github johndoe
```

This will prompt you for:
- Master password (if first time)
- Account password
- Optional notes

### Retrieve a Password

```bash
# Get all usernames for an account
safepass get github

# Get specific username
safepass get github johndoe

# Show password in plaintext
safepass get github johndoe --show
```

### List All Entries

```bash
safepass list
```

## Usage

### Commands Overview

| Command | Description | Example |
|---------|-------------|---------|
| `add` | Add a new password entry | `safepass add github johndoe` |
| `get` | Retrieve account or specific entry | `safepass get github johndoe` |
| `delete` | Delete account or specific usernames | `safepass delete github johndoe` |
| `list` | List all accounts and usernames | `safepass list` |
| `update` | Update password or notes | `safepass update github johndoe` |
| `help` | Get help for any command | `safepass help add` |

### Add Command

Add a new account and username entry to your encrypted vault.

```bash
# Interactive mode (recommended)
safepass add

# With account and username
safepass add github johndoe

# With just account (will prompt for username)
safepass add gmail
```

**What you'll be prompted for:**
- Master password (for vault access)
- Account name (if not provided)
- Username (if not provided)
- Password for the account
- Optional notes

### Get Command

Retrieve details from your vault.

```bash
# Get all usernames for an account
safepass get github

# Get specific username details
safepass get github johndoe

# Show password in plaintext (default is masked)
safepass get github johndoe --show
```

### Delete Command

Remove entries from your vault.

```bash
# Delete entire account and all usernames
safepass delete github

# Delete specific usernames only
safepass delete github johndoe janedoe

# Delete single username
safepass delete gmail alice123
```

### List Command

View all stored accounts and usernames.

```bash
# List with masked passwords (default)
safepass list

# List with plaintext passwords
safepass list --show
```

### Update Command

Modify existing entries.

```bash
safepass update github johndoe
```

**You can update:**
- Password (press Enter to keep existing)
- Notes (press Enter to keep existing)

## Security Features

- **Master Password Protection**: Single master password encrypts your entire vault
- **Local Storage Only**: No cloud dependencies, all data stays on your device
- **Password Masking**: Passwords hidden by default in output
- **Secure Prompts**: Sensitive input handled securely
- **Encrypted Vault File**: All data encrypted at rest

## File Structure

SafePass stores your encrypted vault file locally. The exact location depends on your system configuration.

```
safepass/
├── main.go              # Main application entry
├── cmd/                 # Command implementations
│   ├── add.go
│   ├── get.go
│   ├── delete.go
│   ├── list.go
│   ├── update.go
│   └── root.go
├── internal/            # Internal packages
│   ├── vault/           # Vault operations
│   ├── crypto/          # Encryption/decryption
│   └── utils/           # Utility functions
├── go.mod
├── go.sum
├── README.md
└── LICENSE
```

## Examples

### Basic Workflow

```bash
# Add a GitHub account
safepass add github myusername
# Prompts: master password, account password, optional notes

# Add a Gmail account
safepass add gmail john.doe@gmail.com
# Prompts: master password, account password, optional notes

# List all accounts
safepass list
# Shows: github (myusername), gmail (john.doe@gmail.com)

# Get GitHub password (masked)
safepass get github myusername
# Output: Account: github, Username: myusername, Password: ******

# Get GitHub password (plaintext)
safepass get github myusername --show
# Output: Account: github, Username: myusername, Password: actualpassword

# Update password
safepass update github myusername
# Prompts: master password, new password (optional), new notes (optional)

# Delete specific username
safepass delete github myusername
```

### Multiple Users per Account

```bash
# Add multiple users for the same service
safepass add github personal-account
safepass add github work-account
safepass add github side-project

# List all GitHub accounts
safepass get github
# Shows all three usernames

# Get specific account
safepass get github work-account --show
```

## Development

### Building from Source

```bash
git clone https://github.com/PraneethShetty626/safepass.git
cd safepass
go mod tidy
go build -o safepass
```

### Running Tests

```bash
go test ./...
```

### Cross-Platform Builds

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o safepass-linux

# Windows
GOOS=windows GOARCH=amd64 go build -o safepass.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o safepass-macos
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Security Best Practices

- **Use a Strong Master Password**: This is your only line of defense
- **Keep Your Master Password Safe**: There's no password recovery mechanism
- **Regular Backups**: Back up your vault file regularly
- **Secure Your System**: Ensure your computer is secure since the vault is stored locally
- **Use `--show` Flag Carefully**: Only display passwords when necessary

## FAQ

### Q: Where is my vault file stored?
A: The vault file is stored locally on your system. The exact location depends on your OS and configuration.

### Q: What happens if I forget my master password?
A: Unfortunately, there's no way to recover your data without the master password. This is by design for security.

### Q: Can I use the same account name for different services?
A: Account names should be unique identifiers. Use descriptive names like "github-personal" and "github-work" for different contexts.

### Q: How do I backup my passwords?
A: Locate your vault file and create a secure backup. The file is encrypted, so it's safe to store, but keep your master password separate and secure.

### Q: Can I sync between devices?
A: Currently, SafePass is designed for single-device use. You can manually copy the vault file between devices, but ensure you keep them synchronized.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/PraneethShetty626/safepass/issues)
- **Documentation**: Use `safepass help [command]` for detailed command help
- **Security Issues**: Please report responsibly

---

**⚠️ Important**: Always remember your master password. There is no recovery mechanism by design.