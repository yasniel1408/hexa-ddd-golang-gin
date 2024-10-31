package application

import (
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/auth/domain/valueobjects"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/domain/entities"
)

type AuthService interface {
    Login(credentials valueobjects.Credentials) (string, error)
    Register(user entities.User) error
}