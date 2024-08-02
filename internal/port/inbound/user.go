package inbound

import (
	"context"

	"github.com/marcelofabianov/my-cash/internal/domain"
)

// PKG

type PasswordHasher interface {
	Hash(data string) (string, error)
	Compare(data, encodedHash string) (bool, error)
}

// User / UseCase

type CreateUserUseCaseInput struct {
	Document string
	Name     string
	Email    string
	Password string
}

type CreateUserUseCaseOutput struct {
	User *domain.User
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, input CreateUserUseCaseInput) (*CreateUserUseCaseOutput, error)
}

// User / Service Layer

type CreateUserServiceInput struct {
	CreateUserUseCaseInput
}

type CreateUserServiceOutput struct {
	CreateUserUseCaseOutput
}

type UserService interface {
	CreateUser(ctx context.Context, input CreateUserServiceInput) (*CreateUserServiceOutput, error)
}

// User / Presenter

type UserPresenter struct {
	ID        string `json:"id"`
	Document  string `json:"document"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// User / Request Validate

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Document string `json:"document" validate:"required,min=11,max=14"`
}
