package user

import (
	"errors"
	"regexp"
)

func CheckPasswordPolicy(plainTextPwd string) (bool, error) {
	if plainTextPwd == "" {
		return false, errors.New("패스워드는 필수 값입니다.")
	}
	if len(plainTextPwd) < 8 {
		return false, errors.New("패스워드는 최소 8자 이상입니다.")
	}
	if len(plainTextPwd) > 16 {
		return false, errors.New("패스워드는 최대 16자 입니다.")
	}
	matched, err := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, plainTextPwd)
	if err != nil || !matched {
		return false, errors.New("하나 이상의 특수문자가 포함되어야 합니다.")
	}
	matched, err = regexp.MatchString(`[a-zA-Z]+`, plainTextPwd)
	if err != nil || !matched {
		return false, errors.New("하나 이상의 문자가 포함되어야 합니다.")
	}
	matched, err = regexp.MatchString(`[0-9]+`, plainTextPwd)
	if err != nil || !matched {
		return false, errors.New("하나 이상의 숫자가 포함되어야 합니다.")
	}
	return true, nil
}
