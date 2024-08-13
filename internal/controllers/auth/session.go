package auth

const (
	SessKey string = "session"
	AuthKey string = "Authorization"
)

type Session struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type SessionID struct {
	AccessToken string
}
