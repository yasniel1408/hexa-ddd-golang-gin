package application

import (
	"errors"
	dtos_http "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/input_adapters/http/dtos"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/sql/dao"
	dtos_sql "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/sql/dtos"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/domain/port"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(credentials dtos_http.LoginDto) (string, error)
	Register(user dtos_http.RegisterDto) error
}

type authService struct {
	userRepo port.IUserPort[dtos_sql.RegisterSqlDto, dao.UserDao]
	jwtKey   []byte
}

func AuthService(userRepo port.IUserPort[dtos_sql.RegisterSqlDto, dao.UserDao], jwtKey []byte) IAuthService {
	return &authService{
		userRepo: userRepo,
		jwtKey:   jwtKey,
	}
}

func (s *authService) Register(user dtos_http.RegisterDto) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Create(dtos_sql.RegisterSqlDto(user))
}

func (s *authService) Login(credentials dtos_http.LoginDto) (string, error) {
	user, err := s.userRepo.GetByEmail(credentials.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}