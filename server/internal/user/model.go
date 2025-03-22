package user

import "time"

type User struct {
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Created_at time.Time
	Deleted_at time.Time
}
