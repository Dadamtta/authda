package dadamtta

import (
	"dadamtta/internal/appl"
	"dadamtta/internal/product"
	"dadamtta/internal/user"
	"dadamtta/pkg/apis"
	"dadamtta/pkg/apis/response"
	"dadamtta/pkg/auth"
	"dadamtta/pkg/utils/logger"
	"dadamtta/private/p_appl"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserCommand(router *gin.Engine, repository user.Repository, applRepository appl.Repository, productRepository product.Repository) {
	service := user.NewService(repository, applRepository, productRepository)
	SignUp(router, service)
	SignIn(router, service)
	CreateWeddingApp(router, service)
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

func CreateWeddingApp(router *gin.Engine, service user.Service) {
	router.POST("/v1/users/create/wedding-app", func(c *gin.Context) {
		bearerToken := apis.GetBearerToken(c)
		if len(bearerToken) == 0 {
			logger.Error(fmt.Sprintf("[Token] Empty."))
			// 401
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorResponse(response.ERROR_AUTHORIZED, ""))
			return
		}
		token, err := apis.ParseAccessToken(bearerToken)
		if err != nil {
			logger.Error(fmt.Sprintf("[Token] Parse Error. Error -> %s", err.Error()))
			// 401
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorResponse(response.ERROR_AUTHORIZED, ""))
			return
		}
		dto := &UserAppRegisterRequest{}
		err = apis.BodyMapper(c, dto)
		if err != nil {
			logger.Error("[Mapping] Request DTO Decrypt Mapping Error.")
			// 400
			c.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponse(response.ERROR_DTO_UNMARSHAL, ""))
			return
		}
		appId, err := service.CreateApp(user.WEDDING, token.Id, dto.ProductId)
		if err != nil {
			logger.Error(fmt.Sprintf("[Func] Create App Error. User Id -> %s, Product Id -> %s", token.Id, dto.ProductId))
			if response.HandleResponseErrorWithCustomMessage(c, err, "") {
				return
			} else {
				// 500
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
		// 201
		c.AbortWithStatusJSON(http.StatusCreated, &UserAppRegisterResponse{
			AppId: appId,
		})
	})
}

func UpdateWeddingAppData(router *gin.Engine, service user.Service) {
	router.PUT("/v1/users/apps/:appId", func(c *gin.Context) {
		appId := c.Param("appId")
		bearerToken := apis.GetBearerToken(c)
		if len(bearerToken) == 0 {
			logger.Error(fmt.Sprintf("[Token] Empty."))
			// 401
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorResponse(response.ERROR_AUTHORIZED, ""))
			return
		}
		token, err := apis.ParseAccessToken(bearerToken)
		if err != nil {
			logger.Error(fmt.Sprintf("[Token] Parse Error. Error -> %s", err.Error()))
			// 401
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.NewErrorResponse(response.ERROR_AUTHORIZED, ""))
			return
		}

		err = service.UpdateAppData(user.WEDDING, token.Id, appId, &p_appl.WeddingInvitation{})

		// 200
		c.AbortWithStatus(http.StatusOK)
	})
}
