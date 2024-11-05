package application

import (
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/port"
	cache "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/out_adapters/cache"
	dtos_sql "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/out_adapters/sql/dtos"
	"strconv"
	"time"
)

type IUserService interface {
	GetUserByID(id uint) (entities.User, error)
}

type userService struct {
	userRepo port.IUserPort[dtos_sql.RegisterSqlDto, entities.User]
	cache    cache.ICacheUsersAdapter
}

func UserService(userRepo port.IUserPort[dtos_sql.RegisterSqlDto, entities.User], cache cache.ICacheUsersAdapter) IUserService {
	return &userService{userRepo, cache}
}

func (s *userService) GetUserByID(id uint) (entities.User, error) {
	data, isSaved := s.cache.Get(strconv.Itoa(int(id)))

	if isSaved {
		return data.(entities.User), nil
	}

	newData, _ := s.userRepo.GetByID(id)

	err := s.cache.Set(strconv.Itoa(int(newData.ID)), newData, time.Minute*1)
	if err != nil {
		return entities.User{}, err
	}

	return newData, nil
}
