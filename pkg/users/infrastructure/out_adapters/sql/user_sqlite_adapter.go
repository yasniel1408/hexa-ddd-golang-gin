package sql

import (
	"errors"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/port"
	dtos_sql "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/out_adapters/sql/dtos"

	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) port.IUserPort[dtos_sql.RegisterSqlDto, entities.User] {
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

func (r *userRepository) Create(user dtos_sql.RegisterSqlDto) error {
	return r.db.Create(&user).Error
}
