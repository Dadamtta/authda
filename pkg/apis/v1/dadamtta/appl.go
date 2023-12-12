package dadamtta

import (
	"dadamtta/internal/payment_order"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAppCommand(router *gin.Engine, paymentOrderRepository payment_order.Repository) {
	GetAppData(router)
}

func GetAppData(router *gin.Engine) {
	router.GET("/v1/apps/:appId/data", func(c *gin.Context) {
		appId := c.Param("appId")

		// 만료된 APP인지 확인

		// 결제했는지 확인

		c.AbortWithStatusJSON(http.StatusOK, &UserAppRegisterResponse{
			AppId: appId,
		})
	})
}
