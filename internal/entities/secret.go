package entities

import "time"

type Secret struct {
	Name      string
	Data      string
	User      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
