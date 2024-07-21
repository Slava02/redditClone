package entities

type Session struct {
	UserID   string
	Username string
	Token    string
}

func NewSession() *Session {
	return &Session{}
}
