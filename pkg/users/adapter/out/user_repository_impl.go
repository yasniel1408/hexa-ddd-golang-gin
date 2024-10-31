package out

import (
    "errors"

    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/repositories"
    "gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
    return &userRepository{db}
}

func (r *userRepository) GetByID(id uint) (entities.User, error) {
    var user entities.User
    result := r.db.First(&user, id)
    if result.Error != nil {
        return entities.User{}, result.Error
    }
    return user, nil
}

func (r *userRepository) GetByEmail(email string) (entities.User, error) {
    var user entities.User
    result := r.db.Where("email = ?", email).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return entities.User{}, errors.New("user not found")
        }
        return entities.User{}, result.Error
    }
    return user, nil
}

func (r *userRepository) Create(user entities.User) error {
    return r.db.Create(&user).Error
}