package repositories

import (
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/in/dtos"
)

type UserRepository interface {
	GetByID(id uint) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
	Create(user dtos.RegisterDto) error
}
