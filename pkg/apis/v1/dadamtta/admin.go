package dadamtta

import (
	"dadamtta/internal/admin"
	"dadamtta/pkg/apis"
	"dadamtta/pkg/apis/response"
	"dadamtta/pkg/utils/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAdminCommand(router *gin.Engine, repository admin.Repository) {
	service := admin.NewService(repository)
	Login(router, service)
}

func Login(router *gin.Engine, service admin.Service) {
	router.POST("/v1/admins/login", func(c *gin.Context) {
		dto := AdminLogInFormRequest{}
		err := apis.BodyMapperWithDecrypt(c, &dto)
		if err != nil {
			logger.Error("[Mapping] Request DTO Decrypt Mapping Error.")
			// 400
			c.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponse(response.ERROR_DTO_UNMARSHAL, ""))
			return
		}
		err = service.Login(dto.Id, dto.Pwd)
		if err != nil {
			logger.Error(fmt.Sprintf("[Func] Login Fail. ID -> %s", dto.Id))
			// 401
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorResponse(response.ERROR_AUTHORIZED, ""))
			return
		}
		token, err := apis.GenerateAdminAccessToken(dto.Id)
		if err != nil {
			logger.Error(fmt.Sprintf("[Token] Generate Token Error. ID -> %s", dto.Id))
			// 500
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorResponse(response.ERROR_ENTITY_NOTFOUND, ""))
			return
		}
		// 200
		c.AbortWithStatusJSON(http.StatusOK, AdminTokenResponse{
			AccessToken: token,
		})
	})
}
