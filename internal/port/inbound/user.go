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
