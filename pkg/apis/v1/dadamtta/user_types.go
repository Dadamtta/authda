package dadamtta

type UserRegisterFormRequest struct {
	Id    string `json:"id"`
	Pwd   string `json:"pwd"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// 홈페이지 입력 폼에서 로그인 요청
type UserSignInFormRequest struct {
	Id  string `json:"id"`
	Pwd string `json:"pwd"`
}

type UserTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type UserAppRegisterRequest struct {
	ProductId string `json:"product_id"`
}

type UserAppRegisterResponse struct {
	AppId string `json:"app_id"`
}
