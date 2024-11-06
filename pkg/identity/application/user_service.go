package application

import (
	"fmt"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/domain/port"
	dtos_http "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/input_adapters/http/dtos"
	cache "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/cache"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/sql/dao"
	dtos_sql "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/sql/dtos"
	"strconv"
	"time"
)

type IUserService interface {
	GetUserByID(id uint) (dtos_http.GetUserResponseDto, error)
}

type userService struct {
	userRepo port.IUserPort[dtos_sql.RegisterSqlDto, dao.UserDao]
	cache    cache.ICacheUsersAdapter
}

func UserService(userRepo port.IUserPort[dtos_sql.RegisterSqlDto, dao.UserDao], cache cache.ICacheUsersAdapter) IUserService {
	return &userService{userRepo, cache}
}

func (s *userService) GetUserByID(id uint) (dtos_http.GetUserResponseDto, error) {
	data, isSaved := s.cache.Get(strconv.Itoa(int(id)))

	if isSaved {
		fmt.Print(data)
		return dtos_http.GetUserResponseDto(data.(dao.UserDao)), nil
	}

	newData, _ := s.userRepo.GetByID(id)

	err := s.cache.Set(strconv.Itoa(int(newData.ID)), newData, time.Minute*1)
	if err != nil {
		return dtos_http.GetUserResponseDto(dao.UserDao{}), err
	}

	return dtos_http.GetUserResponseDto(newData), nil
}
