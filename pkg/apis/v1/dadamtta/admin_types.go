package dadamtta

type AdminLogInFormRequest struct {
	Id  string `json:"id"`
	Pwd string `json:"pwd"`
}

type AdminTokenResponse struct {
	AccessToken string `json"access_token"`
}
