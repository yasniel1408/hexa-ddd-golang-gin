package errors

import "errors"

var ErrInvalidEmail = errors.New("invalid email format")

var ErrUserNotFound = errors.New("user not found")

var NameCannotBeEmpty = errors.New("name cannot be empty")
var PasswordCannotBeEmpty = errors.New("password cannot be empty")
