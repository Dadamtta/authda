package dadamtta

import (
	"dadamtta/internal/user"
	"dadamtta/pkg/apis/v1/dadamtta"
	"dadamtta/pkg/utils/logger"
	"dadamtta/private/p_appl"
	"dadamtta/private/p_payment_order"
	"dadamtta/private/p_product"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCommand(router *gin.Engine, db *gorm.DB) {
	logger.Debug("Route SET")
	// route 등록
	dadamtta.NewUserCommand(
		router,
		user.NewRdbRepository(db),
		p_appl.NewApplRdbRepository(db),
		p_appl.NewApplDataNoSqlRepository(db), // db 다른 것으로
		p_product.NewRdbRepository(db),
		p_payment_order.NewPaymentOrderRdbRepository(db),
	)
	dadamtta.NewRSACommand(router)
}
