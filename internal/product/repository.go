package product

import "gorm.io/gorm"

type Repository interface {
	Save(product product) error
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
