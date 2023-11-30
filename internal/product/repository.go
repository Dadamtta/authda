package product

import (
	"dadamtta/internal/sql"

	"gorm.io/gorm"
)

type Repository interface {
	Save(product product) error
	Search(options *sql.SearchOptions)
}

type rdbRepository struct {
	db *gorm.DB
}

func NewRdbRepository(db *gorm.DB) Repository {
	return &rdbRepository{
		db: db,
	}
}

func (*rdbRepository) Save(product product) error {
	return nil
}

func (*rdbRepository) Search(options *sql.SearchOptions) {

}
