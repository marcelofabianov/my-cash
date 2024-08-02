package usecase

import (
	"context"

	"github.com/marcelofabianov/my-cash/internal/domain"
	dError "github.com/marcelofabianov/my-cash/internal/domain/error"
	"github.com/marcelofabianov/my-cash/internal/port/inbound"
	"github.com/marcelofabianov/my-cash/internal/port/outbound"
)

type CreateUserUseCase struct {
	repository outbound.CreateUserRepository
	hasher     inbound.PasswordHasher
}

func NewCreateUserUseCase(repository outbound.CreateUserRepository, hasher inbound.PasswordHasher) *CreateUserUseCase {
	return &CreateUserUseCase{repository: repository, hasher: hasher}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, input inbound.CreateUserUseCaseInput) (*inbound.CreateUserUseCaseOutput, error) {
	hashedPassword, err := u.hasher.Hash(input.Password)
	if err != nil {
		return nil, dError.NewUserPasswordHashFailedError(err)
	}

	user, err := domain.NewUser(input.Document, input.Name, input.Email, hashedPassword)
	if err != nil {
		return nil, dError.NewUserInvalidEntityDataError(err)
	}

	inputRepo := outbound.CreateUserRepositoryInput{
		User: user,
	}

	if err := u.repository.Create(ctx, inputRepo); err != nil {
		return nil, dError.NewUserNotCreatedError(err)
	}

	return &inbound.CreateUserUseCaseOutput{
		User: user,
	}, nil
}
