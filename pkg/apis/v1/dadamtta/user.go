package dadamtta

import (
	"dadamtta/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserCommand(router *gin.Engine, repository user.Repository) {
	service := user.NewService(repository)
	SignUp(router, service)
}

func SignUp(router *gin.Engine, service user.Service) {
	router.POST("/v1/users/sign-up", func(c *gin.Context) {
		// 데이터 RSA 공개키 암호화 처리 되어 있음 (session에서 제어)
		userId, err := service.SignUp("", "", "", "", "", 0, 0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": userId,
			})
		}
	})
}

func SignIn(router *gin.Engine, service user.Service) {
	router.POST("/v1/users/sign-in", func(c *gin.Context) {
		// 아이디, 패스워드 RSA 공개키 암호화 처리 되어 있음 (session에서 제어)
	})
}
