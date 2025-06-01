package vault

import "time"

type Entry struct {
	Password  string
	UpdatedAt time.Time
	Note      string
}
