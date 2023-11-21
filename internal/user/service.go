package user

// 회원가입
// 로그인
type service struct {
}

func NewService() *service {
	return &service{}
}

func (*service) SignUp() {
	println("회원가입 진행")
}

func (*service) SignIn() {
	println("로그인 진행")
}
