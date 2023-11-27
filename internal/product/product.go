package product

import "errors"

type product struct {
	Id           string
	CategoryCode string
	Label        string
	Price        uint32
	Description  string
	Content      string
	Hits         uint32
	State        uint8
}

func GenerateProduct(categoryCode, label string, price uint32, description, content string) (*product, error) {
	if label == "" {
		return nil, errors.New("상품명 정보 입력 누락")
	}
	// todo category code 확인

	return &product{
		Id:           "",
		CategoryCode: "",
		Label:        "",
		Price:        0,
		Description:  "",
		Content:      "",
		Hits:         0,
		State:        0,
	}, nil
}
