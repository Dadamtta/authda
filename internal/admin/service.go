package admin

import (
	"dadamtta/internal/common/errorc"
	"dadamtta/pkg/utils/logger"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(id, plainPwd string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Login(id, plainPwd string) (err error) {
	admin := s.repository.FindById(id)
	if admin == nil {
		err = errorc.EntityNotFoundError
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.HashedPwd), []byte(plainPwd))
	if err != nil {
		logger.Error(fmt.Sprintf("[Token] Generate Token Error. ID -> %s", id))
	}
	return
}
