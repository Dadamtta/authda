package product

import (
	"dadamtta/internal/sql"
)

type Repository interface {
	Save(product Product) error
	Search(options *sql.SearchOptions)
	FindById(id string) *Product
}

type CategoryRepository interface {
	Save(category Category) error
	FindByCode(code string) *Category
	UpdateByCode(code, name string) *Category
	DeleteByCode(code string) error
}
