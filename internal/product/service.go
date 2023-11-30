package product

import "dadamtta/internal/sql"

type Service interface {
	Register(adminId, categoryCode, label string, price uint32, description, content string) (string, error)
	Search(options *sql.SearchOptions)
	Get(productId string)
	Update()
	Delete(productId string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Register(adminId, categoryCode, label string, price uint32, description, content string) (id string, err error) {
	// admin

	// 카테고리 확인 todo 카테고리 조회
	product, err := GenerateProduct(adminId, categoryCode, label, price, description, content)
	if err != nil {
		return
	}
	err = s.repository.Save(*product)
	if err != nil {
		return
	}
	id = product.Id
	return
}

func (s *service) Search(options *sql.SearchOptions) {

}

func (s *service) Get(productId string) {

}

func (s *service) Update() {

}

func (s *service) Delete(productId string) {

}
