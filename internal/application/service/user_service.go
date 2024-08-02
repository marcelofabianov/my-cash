package service

import (
	"context"

	"github.com/marcelofabianov/my-cash/internal/port/inbound"
	"github.com/marcelofabianov/my-cash/pkg/logger"
)

type UserService struct {
	log        logger.Logger
	createUser inbound.CreateUserUseCase
}

func NewUserService(log logger.Logger, createUser inbound.CreateUserUseCase) *UserService {
	return &UserService{log, createUser}
}

func (s *UserService) CreateUser(ctx context.Context, input inbound.CreateUserServiceInput) (*inbound.CreateUserServiceOutput, error) {
	output, err := s.createUser.Execute(ctx, input.CreateUserUseCaseInput)
	if err != nil {
		return nil, err
	}

	// Todo: dispatch user.created event

	s.log.Info("user created", s.log.Field("user", output.User))

	return &inbound.CreateUserServiceOutput{
		CreateUserUseCaseOutput: *output,
	}, nil
}
