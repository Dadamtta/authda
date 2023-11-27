package dadamtta

import (
	"dadamtta/internal/user"
	"dadamtta/pkg/apis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserCommand(router *gin.Engine, repository user.Repository) {
	service := user.NewService(repository)
	SignUp(router, service)
}

func SignUp(router *gin.Engine, service user.Service) {
	router.POST("/v1/users/sign-up", func(c *gin.Context) { // Body RSA Encoded
		dto := UserRegisterFormRequest{}
		apis.BodyMapperWithDecrypt(c, &dto)
		userId, err := service.SignUp(dto.Id, dto.Pwd, dto.Phone, dto.Email, dto.Name, dto.Age, dto.Gender)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": userId,
			})
		}
		return
	})

}

func SignIn(router *gin.Engine, service user.Service) {
	router.POST("/v1/users/sign-in", func(c *gin.Context) { // Body RSA Encoded
		dto := UserSignInFormRequest{}
		apis.BodyMapperWithDecrypt(c, &dto)
		err := service.SignIn(dto.Id, dto.Pwd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else {
			c.Status(http.StatusNoContent)
		}
		return
	})
}
