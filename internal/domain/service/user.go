package service

type UserRepository interface {
	Login() error
	Signup() (string, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Login() error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Signup() (string, error) {
	//TODO implement me
	panic("implement me")
}
