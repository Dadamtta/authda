package apis

import (
	"dadamtta/pkg/auth"
	"dadamtta/private/p_policy"

	"github.com/gin-gonic/gin"
)

type Token struct {
	Id        string
	Authority string
}

func GetBearerToken(c *gin.Context) (bearerToken string) {
	bearerToken = c.Request.Header.Get("Authorization")
	return
}

func GenerateAdminAccessToken(adminId string) (token string, err error) {
	token, err = auth.New("access-token").GenerateToken(map[string]any{
		"Id":        adminId,
		"Authority": "ADMIN",
		"ExpiresAt": p_policy.ACCESS_TOKEN_EXPIRED_AT,
	})
	return
}

func GenerateUserAccessToken(userId string) (token string, err error) {
	token, err = auth.New("access-token").GenerateToken(map[string]any{
		"Id":        userId,
		"Authority": "USER",
		"ExpiresAt": p_policy.ACCESS_TOKEN_EXPIRED_AT,
	})
	return
}

func ParseAccessToken(accessToken string) (*Token, error) {
	claims, err := auth.New("access-token").Parse(accessToken)
	if err != nil {
		return nil, err
	}
	return &Token{
		Id:        claims["Id"].(string),
		Authority: claims["Authority"].(string),
	}, nil
}
