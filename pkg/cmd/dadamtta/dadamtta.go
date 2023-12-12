package dadamtta

import (
	"dadamtta/internal/product"
	"dadamtta/internal/user"
	"dadamtta/pkg/apis/v1/dadamtta"
	"dadamtta/pkg/utils/logger"
	"dadamtta/private/p_appl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCommand(router *gin.Engine, db *gorm.DB) {
	logger.Debug("Route SET")
	// route 등록
	dadamtta.NewUserCommand(router, user.NewRdbRepository(db), p_appl.NewApplRdbRepository(db), product.NewRdbRepository(db))
	dadamtta.NewRSACommand(router)
}
