package card

import "time"

type Card struct {
	ID           int
	Title        string
	Content      string
	Author       string
	CreatedAt    time.Time
	CreatedAtStr string
}
