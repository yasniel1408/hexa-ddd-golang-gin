package application

import (
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/repositories"
)

type UserService interface {
	GetUserByID(id uint) (entities.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetUserByID(id uint) (entities.User, error) {
	return s.userRepo.GetByID(id)
}
