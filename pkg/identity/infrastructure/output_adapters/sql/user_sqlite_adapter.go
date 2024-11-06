package sql

import (
	"errors"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/domain/port"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/sql/dao"
	dtos_sql "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/sql/dtos"

	"gorm.io/gorm"
)

type userSqliteAdapter struct {
	db *gorm.DB
}

func UserSqliteAdapter(db *gorm.DB) port.IUserPort[dtos_sql.RegisterSqlDto, dao.UserDao] {
	return &userSqliteAdapter{db}
}

func (r *userSqliteAdapter) GetByID(id uint) (dao.UserDao, error) {
	var user dao.UserDao
	result := r.db.First(&user, id)
	if result.Error != nil {
		return dao.UserDao{}, result.Error
	}
	return user, nil
}

func (r *userSqliteAdapter) GetByEmail(email string) (dao.UserDao, error) {
	var user dao.UserDao
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return dao.UserDao{}, errors.New("user not found")
		}
		return dao.UserDao{}, result.Error
	}
	return user, nil
}

func (r *userSqliteAdapter) Create(user dtos_sql.RegisterSqlDto) error {
	return r.db.Create(&user).Error
}
