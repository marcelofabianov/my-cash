package outbound

import (
	"context"

	"github.com/marcelofabianov/my-cash/internal/domain"
)

// User / Repository

type CreateUserRepositoryInput struct {
	User *domain.User
}

type CreateUserRepository interface {
	Create(ctx context.Context, input CreateUserRepositoryInput) error
}

type UserRepository interface {
	CreateUserRepository
}
