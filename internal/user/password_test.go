package user

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestCheckPasswordPolicy(t *testing.T) {
	isValid, err := CheckPasswordPolicy("Aad@dkas2sdfasdf")
	if err != nil {
		println("에러가 발생했습니다.")
		println(err.Error())
	} else {
		println(isValid)
	}
}

func TestBcryptPassword(t *testing.T) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte("funch12#$"), 10)
	if err != nil {
		println("에러가 발생했습니다.")
		println(err.Error())
	} else {
		println(string(hashedPw))
	}
}
