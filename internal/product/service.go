package product

type Service interface {
	Register(categoryCode, label string, price uint32, description, content string) (string, error)
	Search()
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

func (s *service) Register(categoryCode, label string, price uint32, description, content string) (id string, err error) {
	// 카테고리 확인 todo 카테고리 조회
	product, err := GenerateProduct(categoryCode, label, price, description, content)
	if err != nil {
		return
	}
	s.repository.Save(*product)
	if err != nil {
		return
	}
	id = product.Id
	return
}

func (s *service) Search() {

}

func (s *service) Get(productId string) {

}

func (s *service) Update() {

}

func (s *service) Delete(productId string) {

}
