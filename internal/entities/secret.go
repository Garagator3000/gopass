package entities

import "time"

type Secret struct {
	Name      string
	Data      string
	User      string
	Group     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
