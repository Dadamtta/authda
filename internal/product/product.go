package product

import (
	"errors"

	"github.com/google/uuid"
)

type State int

const (
	Pause  State = 0
	OnSale State = 1
)

type product struct {
	Id           string
	CategoryCode string
	AdminId      string
	Label        string
	Price        uint32
	Description  string
	Content      string
	Hits         uint32
	State        State
}

func GenerateProduct(adminId, categoryCode, label string, price uint32, description, content string) (*product, error) {
	if label == "" {
		return nil, errors.New("상품명 정보 입력 누락")
	}
	// todo category code 확인
	return &product{
		Id:           uuid.New().String(),
		CategoryCode: categoryCode,
		AdminId:      adminId,
		Label:        label,
		Price:        price,
		Description:  description,
		Content:      content,
		Hits:         0,
		State:        OnSale,
	}, nil
}
