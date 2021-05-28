package user

import (
	domain "github.com/ztjustin/questions_api/domain/users"
	"github.com/ztjustin/questions_api/repository/user"
)

type Service interface {
	GetAll() ([]*domain.User, error)
}

type service struct {
	userRepo user.RestUserRepository
}

func NewService(usersRepo user.RestUserRepository) Service {
	return &service{
		userRepo: usersRepo,
	}
}

func (s *service) GetAll() ([]*domain.User, error) {
	listUser, err := s.userRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return listUser, nil

}
