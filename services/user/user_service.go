package user

import (
	domain "github.com/ztjustin/questions_api/domain/users"
	"github.com/ztjustin/questions_api/repository/user"
)

type Service interface {
	GetAll() ([]*domain.User, error)
	FindById(string) (*domain.User, error)
	Create(*domain.User) (*domain.User, error)
}

type service struct {
	userRepo user.UserRepository
}

func NewService(usersRepo user.UserRepository) Service {
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

func (s *service) FindById(id string) (*domain.User, error) {
	userById, err := s.userRepo.FindById(id)

	if err != nil {
		return nil, err
	}

	return userById, nil
}

func (s *service) Create(newUser *domain.User) (*domain.User, error) {
	user, err := s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}
