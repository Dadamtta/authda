package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type jwToken struct {
	key string
}

func New(key string) *jwToken {
	return &jwToken{
		key,
	}
}

func (j *jwToken) GenerateToken(claims map[string]any) (string, error) {
	claimMap := jwt.MapClaims{}
	for k, v := range claims {
		claimMap[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimMap)
	return token.SignedString([]byte(j.key))
}

func (j *jwToken) IsValid(jwtString string) bool {
	_, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.key), nil
	})
	if err != nil {
		return false
	}
	return true
}

func (j *jwToken) Parse(jwtString string) (map[string]any, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwtString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.key), nil
	})
	return claims, err
}
