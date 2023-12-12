package product

import (
	"dadamtta/internal/sql"

	"gorm.io/gorm"
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

type rdbRepository struct {
	db *gorm.DB
}

func NewRdbRepository(db *gorm.DB) Repository {
	return &rdbRepository{
		db: db,
	}
}

func (*rdbRepository) Save(product Product) error {
	return nil
}

func (*rdbRepository) Search(options *sql.SearchOptions) {

}

func (*rdbRepository) FindById(id string) *Product {
	return nil
}
