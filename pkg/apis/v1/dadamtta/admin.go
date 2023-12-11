package dadamtta

import (
	"dadamtta/internal/admin"
	"dadamtta/pkg/apis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAdminCommand(router *gin.Engine, repository admin.Repository) {
	service := admin.NewService(repository)
	LogIn(router, service)
}

func LogIn(router *gin.Engine, service admin.Service) {
	router.POST("/v1/admins/log-in", func(c *gin.Context) {
		dto := AdminLogInFormRequest{}
		err := apis.BodyMapperWithDecrypt(c, &dto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		err = service.LogIn(dto.Id, dto.Pwd)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{})
			return
		}
		token, err := GenerateAdminAccessToken(dto.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusOK, AdminTokenResponse{
			AccessToken: token,
		})
	})
}
