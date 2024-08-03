package error

import (
	"errors"
)

var (
	ErrUserNotCreated         = errors.New("error_user_not_created")
	ErrUserPasswordHashFailed = errors.New("error_user_password_hash_failed")
	ErrUserExists             = errors.New("error_user_exists")
	ErrUserInvalidEntityData  = errors.New("error_user_invalid_entity_data")
)

func NewUserNotCreatedError(err error) error {
	l := InitLogger()
	l.Error("UserError", l.FieldError(err))

	return ErrUserNotCreated
}

func NewUserPasswordHashFailedError(err error) error {
	l := InitLogger()
	l.Error("UserError", l.FieldError(err))

	return ErrUserPasswordHashFailed
}

func NewUserExistsError(err error) error {
	l := InitLogger()
	l.Error("UserError: User already exists", l.FieldError(err))

	return ErrUserExists
}

func NewUserInvalidEntityDataError(err error) error {
	l := InitLogger()
	l.Error("UserError: Invalid entity data", l.FieldError(err))

	return ErrUserInvalidEntityData
}

func IsUserNotCreatedError(err error) bool {
	return errors.Is(err, ErrUserNotCreated)
}

func IsUserPasswordHashFailedError(err error) bool {
	return errors.Is(err, ErrUserPasswordHashFailed)
}

func IsUserExistsError(err error) bool {
	return errors.Is(err, ErrUserExists)
}

func IsUserInvalidEntityDataError(err error) bool {
	return errors.Is(err, ErrUserInvalidEntityData)
}
