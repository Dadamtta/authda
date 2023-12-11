package admin

type Repository interface {
	FindById(id string) *Admin
}
