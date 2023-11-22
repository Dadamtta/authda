package user

import "gorm.io/gorm"

type Repository interface {
	Save(user user) error
}

type rdbRepository struct {
	db *gorm.DB
}

func NewRdbRepository(db *gorm.DB) Repository {
	return &rdbRepository{
		db: db,
	}
}

func (*rdbRepository) Save(user user) error {
	return nil
}
