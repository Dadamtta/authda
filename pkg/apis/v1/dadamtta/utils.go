package dadamtta

import (
	"dadamtta/pkg/auth"
	"dadamtta/private/p_policy"
)

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
