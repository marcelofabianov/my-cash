package outbound

import (
	"context"

	"github.com/marcelofabianov/my-cash/internal/domain"
)

// User / Repository

type CreateUserRepositoryInput struct {
	User *domain.User
}

type UserExistsRepositoryInput struct {
	Email    string
	Document string
}

type CreateUserRepository interface {
	UserExists(ctx context.Context, input UserExistsRepositoryInput) (bool, error)
	Create(ctx context.Context, input CreateUserRepositoryInput) error
}

type UserRepository interface {
	CreateUserRepository
}
