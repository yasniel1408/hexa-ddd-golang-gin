package application

import (
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/in/dtos"
)

type AuthService interface {
	Login(credentials dtos.LoginDto) (string, error)
	Register(user dtos.RegisterDto) error
}
