package user

type Service interface {
	SignUp(id, pwd, phone, email, name string, age, gender uint8) (string, error)
	SignIn()
}

type service struct {
	repository Repository
}

func NewService(userRepository Repository) Service {
	return &service{
		repository: userRepository,
	}
}

func (s *service) SignUp(id, pwd, phone, email, name string, age, gender uint8) (string, error) {
	println("회원가입 진행")
	newUser, err := GenerateUser(id, pwd, phone, email, name, age, gender)
	if err != nil {
		return "", err
	}
	// todo - Save 에러
	err = s.repository.Save(*newUser)
	if err != nil {
		return "", err
	}
	return newUser.Id, err
}

func (*service) SignIn() {
	println("로그인 진행")
}
