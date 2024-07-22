package entities

import "time"

type User struct {
	ID           int
	Username     string
	Password     string
	RegisteredAt time.Time
}

func NewUser(id int, username, password string, registrationTime time.Time) *User {
	return &User{
		ID:           id,
		Username:     username,
		Password:     password,
		RegisteredAt: registrationTime,
	}
}
