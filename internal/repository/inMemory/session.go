package inMemory

import "time"

type Session struct {
	ExpiresAt time.Time `json:"expiresAt"`
}
