package entities

import "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/domain/value_objects"

type UserEntity struct {
	ID       uint                  `json:"id"`
	Name     string                `json:"name"`
	Email    value_objects.EmailVo `json:"email"`
	Password string                `json:"password"`
	Role     string                `json:"role"`
}
