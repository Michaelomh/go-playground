package db

import (
	"time"
)

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
