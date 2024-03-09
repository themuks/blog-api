package models

import "time"

type User struct {
	ID           int
	Username     string
	Name         string
	Surname      string
	PasswordHash string
	Salt         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
