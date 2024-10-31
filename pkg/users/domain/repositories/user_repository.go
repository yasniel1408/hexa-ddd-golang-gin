package repositories

import "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"

type UserRepository interface {
    GetByID(id uint) (entities.User, error)
    GetByEmail(email string) (entities.User, error)
    Create(user entities.User) error
}