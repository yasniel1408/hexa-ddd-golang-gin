package value_objects

import (
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/domain/errors"
	"regexp"
)

type EmailVo string

func CreateEmail(email string) (EmailVo, error) {
	if !isValidEmail(email) {
		return "", errors.ErrInvalidEmail
	}
	return EmailVo(email), nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
