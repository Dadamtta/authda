package dadamtta

import (
	"dadamtta/internal/admin"
	"dadamtta/internal/common/errorc"
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
			if response.HandleResponseErrorWithCustomMessage(c, errorc.DtoUnmarshalError, "") {
				return
			} else {
				c.Status(http.StatusBadRequest)
				return
			}
		}
		err = service.Login(dto.Id, dto.Pwd)
		if err != nil {
			logger.Error(fmt.Sprintf("[Func] Login Fail. ID -> %s", dto.Id))
			// 401
			if response.HandleResponseErrorWithCustomMessage(c, errorc.AuthorizedError, "") {
				return
			} else {
				c.Status(http.StatusUnauthorized)
				return
			}
		}
		token, err := GenerateAdminAccessToken(dto.Id)
		if err != nil {
			logger.Error(fmt.Sprintf("[Token] Generate Token Error. ID -> %s", dto.Id))
			// 500
			if response.HandleResponseErrorWithCustomMessage(c, errorc.TokenGenerateError, "") {
				return
			} else {
				c.Status(http.StatusInternalServerError)
				return
			}
		}
		// 200
		c.AbortWithStatusJSON(http.StatusOK, AdminTokenResponse{
			AccessToken: token,
		})
	})
}
