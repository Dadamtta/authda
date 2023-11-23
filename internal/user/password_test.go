package user

import "testing"

func TestCheckPasswordPolicy(t *testing.T) {
	isValid, err := CheckPasswordPolicy("Aad@dkas2sdfasdf")
	if err != nil {
		println("에러가 발생했습니다.")
		println(err.Error())
	} else {
		println(isValid)
	}
}
