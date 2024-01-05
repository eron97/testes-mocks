package service

type UserService interface {
	WelcomeHandler() string
	RegisterHandler(name string) bool
}

type UserServiceImpl struct {
	users []string
}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{
		users: []string{},
	}
}

func (s *UserServiceImpl) Welcome() string {
	return "Bem-vindo ao servidor Gin!"
}

func (s *UserServiceImpl) Register(name string) bool {
	s.users = append(s.users, name)
	return true
}
