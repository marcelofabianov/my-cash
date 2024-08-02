package error

import (
	"errors"
)

var (
	ErrUserNotCreated            = errors.New("error_user_not_created")
	ErrUserPasswordHashFailed    = errors.New("error_user_password_hash_failed")
	ErrUserEmailAlreadyExists    = errors.New("error_user_email_already_exists")
	ErrUserDocumentAlreadyExists = errors.New("error_user_document_already_exists")
	ErrUserInvalidEntityData     = errors.New("error_user_invalid_entity_data")
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

func NewUserEmailAlreadyExistsError() error {
	l := InitLogger()
	l.Error("UserError: Email already exists")

	return ErrUserEmailAlreadyExists
}

func NewUserDocumentAlreadyExistsError() error {
	l := InitLogger()
	l.Error("UserError: Document already exists")

	return ErrUserDocumentAlreadyExists
}

func NewUserInvalidEntityDataError(err error) error {
	l := InitLogger()
	l.Error("UserError: Invalid entity data", l.FieldError(err))

	return ErrUserInvalidEntityData
}
