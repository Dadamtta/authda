package p_payment_order

import (
	"dadamtta/internal/payment_order"

	"gorm.io/gorm"
)

type rdbPaymentOrderRepository struct {
	db *gorm.DB
}

func NewPaymentOrderRdbRepository(db *gorm.DB) payment_order.Repository {
	return &rdbPaymentOrderRepository{
		db: db,
	}
}

func (p *rdbPaymentOrderRepository) ExistsByAppId(appId string) bool {
	return true
}
