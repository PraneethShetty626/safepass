package vault

import (
	"encoding/json"
	"errors"
	"os"
)

type Vault struct {
	Entries map[string]PEntry
}

func NewVault() *Vault {
	return &Vault{Entries: make(map[string]PEntry)}
}

func (v *Vault) AddEntry(account string, username string, entry Entry) error {

	vault := *v

	pentry, ok := vault.Entries[account]

	if !ok {
		pentry = make(PEntry)
	}

	err := pentry.addPEntry(username, entry)

	if err != nil {
		return err
	}

	vault.Entries[account] = pentry

	return nil

}

func (v *Vault) DeleteEntry(account string) error {
	vault := *v

	_, err := vault.GetEntry(account)

	if err != nil {
		return err
	}

	delete(v.Entries, account)

	return nil
}

func (v *Vault) GetEntry(account string) (PEntry, error) {
	vault := *v

	pEntry, ok := vault.Entries[account]

	if !ok {
		return nil, errors.New("no account exist")
	}

	return pEntry, nil
}

func (v *Vault) SaveVault(filename, masterPassword string) error {
	data, err := json.Marshal(v.Entries)
	if err != nil {
		return err
	}

	key := deriveKey(masterPassword)
	encrypted, err := encrypt(data, key)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, encrypted, 0600)
}

func LoadVault(filename, masterPassword string) (*Vault, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return NewVault(), nil
		}
		return nil, err
	}

	key := deriveKey(masterPassword)
	decrypted, err := decrypt(data, key)
	if err != nil {
		return nil, errors.New("incorrect password or corrupted vault")
	}

	var entries map[string]PEntry
	if err := json.Unmarshal(decrypted, &entries); err != nil {
		return nil, err
	}

	return &Vault{Entries: entries}, nil
}
