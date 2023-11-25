package product

type Service interface {
	Register(categoryCode, label string, price uint32, description, content string)
	Search()
	Get(productId string)
	Update()
	Delete(productId string)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Register(categoryCode, label string, price uint32, description, content string) {

}

func (s *service) Search() {

}

func (s *service) Get(productId string) {

}

func (s *service) Update() {

}

func (s *service) Delete(productId string) {

}
