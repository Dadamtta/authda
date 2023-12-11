package admin

type Service interface {
	LogIn(id, plainPwd string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) LogIn(id, plainPwd string) error {
	admin := s.repository.FindById(id)

	println(admin)

	return nil
}
