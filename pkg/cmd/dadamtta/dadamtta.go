package dadamtta

import (
	"dadamtta/internal/user"
	"dadamtta/pkg/apis/v1/dadamtta"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCommand(router *gin.Engine, db *gorm.DB) {
	// route 등록
	dadamtta.NewUserCommand(router, user.NewRdbRepository(db))
	dadamtta.NewRSACommand(router)
}
