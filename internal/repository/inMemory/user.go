package inMemory

type Users struct{}

func (u *Users) Login() error {
	//TODO implement me
	panic("implement me")
}

func (u *Users) Signup() (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewUsers() *Users {
	return &Users{}
}
