package dadamtta

type UserRegisterFormRequest struct {
	Id         string `json:"id"`
	EncodedPwd string `json:"encoded_pwd"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Age        uint8  `json:"age"`
	Gender     uint8  `json:"gender"`
}
