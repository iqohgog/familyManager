package user

import "time"

type User struct {
	FirstName  string
	LastName   string
	Email      string
	HashPass   string
	CreatedAt time.Time
	DeletedAt time.Time
}
