package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/marcelofabianov/my-cash/internal/port/outbound"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) UserExists(ctx context.Context, input outbound.UserExistsRepositoryInput) (bool, error) {
	args := m.Called(ctx, input)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) Create(ctx context.Context, input outbound.CreateUserRepositoryInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}
