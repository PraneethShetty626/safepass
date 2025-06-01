package vault

import (
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aquasecurity/table"
)

type PEntry map[string]Entry

func (p *PEntry) addPEntry(username string, entry Entry) error {
	pEntry := *p

	_, ok := pEntry[username]

	if ok {
		return errors.New("username already existes, please use update")
	}

	pEntry[username] = entry

	return nil

}

func (p *PEntry) DeletePEntry(username string) error {
	pEntry := *p

	_, ok := pEntry[username]

	if !ok {
		return errors.New("username doesnt exists")
	}

	delete(pEntry, username)

	return nil
}

func (p *PEntry) GetEntry(username string) (Entry, error) {
	pEntry := *p

	entry, ok := pEntry[username]

	if !ok {
		return Entry{}, errors.New("username doesnt exists")
	}

	return entry, nil
}

func (p *PEntry) Print(showPassword bool) {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Username", "Password", "Note", "Updated At")

	// Collect entries into a slice
	var entries []struct {
		Username string
		Entry    Entry
	}
	for username, entry := range *p {
		entries = append(entries, struct {
			Username string
			Entry    Entry
		}{
			Username: username,
			Entry:    entry,
		})
	}

	// Sort entries by UpdatedAt in descending order
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Entry.UpdatedAt.After(entries[j].Entry.UpdatedAt)
	})

	// Add sorted rows to table
	for i, item := range entries {
		password := item.Entry.Password
		if !showPassword {
			password = strings.Repeat("*", len(password))
		}

		table.AddRow(
			strconv.Itoa(i+1),
			item.Username,
			password,
			item.Entry.Note,
			item.Entry.UpdatedAt.Format(time.RFC1123),
		)
	}

	table.Render()
}
