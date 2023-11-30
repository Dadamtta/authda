package dadamtta

import (
	"dadamtta/internal/user"
	"dadamtta/pkg/apis"
	"dadamtta/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserCommand(router *gin.Engine, repository user.Repository) {
	service := user.NewService(repository)
	SignUp(router, service)
	SignIn(router, service)
}

func SignUp(router *gin.Engine, service user.Service) {
	router.POST("/v1/users/sign-up", func(c *gin.Context) { // Body RSA Encoded
		dto := UserRegisterFormRequest{}
		apis.BodyMapperWithDecrypt(c, &dto)
		userId, err := service.SignUp(dto.Id, dto.Pwd, dto.Phone, dto.Email, dto.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": userId,
			})
			return
		}
	})
}

func SignIn(router *gin.Engine, service user.Service) {
	router.POST("/v1/users/sign-in", func(c *gin.Context) { // Body RSA Encoded
		dto := UserSignInFormRequest{}
		err := apis.BodyMapperWithDecrypt(c, &dto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = service.SignIn(dto.Id, dto.Pwd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			token, err := auth.New("access-token").GenerateToken(map[string]any{
				"ExpiresAt": 60 * 60 * 24,
			})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, UserTokenResponse{
				AccessToken: token,
			})
			return
		}
	})
}
