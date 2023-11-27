package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Id        string
	HashedPwd string
	Email     string
	Name      string
	Phone     string
	Age       uint8
	Gender    uint8
	Deleted   bool
}

func NewUser() *user {
	return &user{}
}

func GenerateUser(id, pwd, phone, email, name string, age, gender uint8) (*user, error) {
	if id == "" {
		return nil, errors.New("ID 정보는 필수 값 입니다")
	}
	// 로그인 요청한 대상을 식별하기 위한 정보
	if phone == "" {
		return nil, errors.New("휴대폰 정보는 필수 값 입니다")
	}
	if pwd == "" {
		return nil, errors.New("회원가입 시 패스워드 필수 입력")
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return nil, errors.New("패스워드 해싱 실패")
	}
	return &user{
		Id:        id,
		HashedPwd: string(hashedPwd),
		Name:      name,
		Phone:     phone,
		Age:       age,
		Gender:    gender,
		Deleted:   false,
	}, nil
}
