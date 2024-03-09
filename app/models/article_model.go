package models

import "time"

type Article struct {
	ID          int
	Title       string
	Description string
	Text        string
	Author      User
	Comments    []Comment
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
