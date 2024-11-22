package domain

import (
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/domain/entities"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/domain/errors"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/domain/value_objects"
)

type UserFactory struct{}

func (f *UserFactory) NewUser(id uint, name string, email string, password string, role string) (*entities.UserEntity, error) {
	emailVo, err := value_objects.CreateEmail(email)
	if err != nil {
		return nil, err
	}

	if name == "" {
		return nil, errors.NameCannotBeEmpty
	}

	if password == "" {
		return nil, errors.PasswordCannotBeEmpty
	}

	return &entities.UserEntity{
		ID:       id,
		Name:     name,
		Email:    emailVo,
		Password: password,
		Role:     role,
	}, nil
}
