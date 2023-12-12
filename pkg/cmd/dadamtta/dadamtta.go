package dadamtta

import (
	"dadamtta/internal/user"
	"dadamtta/pkg/apis/v1/dadamtta"
	"dadamtta/pkg/utils/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCommand(router *gin.Engine, db *gorm.DB) {
	logger.Debug("Route SET")
	// route 등록
	dadamtta.NewUserCommand(router, user.NewRdbRepository(db))
	dadamtta.NewRSACommand(router)
}
