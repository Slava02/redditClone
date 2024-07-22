package entities

import "time"

type Session struct {
	ExpiresAt time.Time `json:"expiresAt"`
}

func NewSession() *Session {
	return &Session{}
}
