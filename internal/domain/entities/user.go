package entities

// pure business logic
type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// business logic with implementation logic
type UserExtend struct {
	User
	ID string
}

func NewUserExtend(user User, id string) UserExtend {
	return UserExtend{
		User: user,
		ID:   id,
	}
}
