package product

import (
	"errors"
	"time"

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
		State:        OnSale,
	}, nil
}

type Category struct {
	Code       string
	Name       string
	ParentCode string
	Layer      uint8
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func GenerateCategory(code string, name string, parentCode string, layer uint8) *Category {
	return &Category{
		Code:       code,
		Name:       name,
		ParentCode: parentCode,
		Layer:      layer,
		CreatedAt:  time.Now(),
	}
}
