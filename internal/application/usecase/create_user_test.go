package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/marcelofabianov/my-cash/internal/application/usecase"
	"github.com/marcelofabianov/my-cash/internal/domain"
	"github.com/marcelofabianov/my-cash/internal/port/inbound"
	inboundMock "github.com/marcelofabianov/my-cash/internal/port/inbound/mock"
	"github.com/marcelofabianov/my-cash/internal/port/outbound"
	outboundMock "github.com/marcelofabianov/my-cash/internal/port/outbound/mock"
)

type CreateUserUseCaseTestSuite struct {
	suite.Suite
	usecase *usecase.CreateUserUseCase
	repo    *outboundMock.MockUserRepository
	hasher  *inboundMock.MockPasswordHasher
}

func (s *CreateUserUseCaseTestSuite) SetupTest() {
	s.repo = new(outboundMock.MockUserRepository)
	s.hasher = new(inboundMock.MockPasswordHasher)
	s.usecase = usecase.NewCreateUserUseCase(s.repo, s.hasher)
}

func (s *CreateUserUseCaseTestSuite) TearDownTest() {
	s.repo.AssertExpectations(s.T())
	s.hasher.AssertExpectations(s.T())
}

func (s *CreateUserUseCaseTestSuite) TestExecute_Success() {
	// Arrange
	inputUC := inbound.CreateUserUseCaseInput{
		Document: "12345678901",
		Name:     "Marcelo",
		Email:    "marcelo@email.com",
		Password: "plain-password",
	}

	hashedPassword := "hashed-password"
	ctx := context.Background()

	// Mock expectations

	s.hasher.On("Hash", inputUC.Password).Return(hashedPassword, nil)

	s.repo.On("UserExists", ctx, mock.MatchedBy(func(args outbound.UserExistsRepositoryInput) bool {
		return args.Document == inputUC.Document || args.Email == inputUC.Email
	})).Return(false, nil)

	s.repo.On("Create", ctx, mock.MatchedBy(func(args outbound.CreateUserRepositoryInput) bool {
		return args.User.Document == domain.Document(inputUC.Document) &&
			args.User.Name == inputUC.Name && args.User.Email == domain.Email(inputUC.Email) && args.User.Password == domain.Password(hashedPassword)
	})).Return(nil)

	// Act
	outputUC, err := s.usecase.Execute(ctx, inputUC)

	// Assert
	s.NoError(err)
	s.NotNil(outputUC)
	s.Equal(outputUC.User.Document.String(), inputUC.Document)
	s.Equal(outputUC.User.Name, inputUC.Name)
	s.Equal(outputUC.User.Email.String(), inputUC.Email)
	s.NotEmpty(outputUC.User.ID.String())
	s.NotEmpty(outputUC.User.CreatedAt.String())
	s.NotEmpty(outputUC.User.UpdatedAt.String())
	s.Nil(outputUC.User.DeletedAt)
	s.False(outputUC.User.Enabled.Bool())
}

func (s *CreateUserUseCaseTestSuite) TestExecute_UserExistsError() {
	//...
}

func (s *CreateUserUseCaseTestSuite) TestExecute_PasswordHashFailedError() {
	//...
}

func (s *CreateUserUseCaseTestSuite) TestExecute_UserInvalidEntityDataError() {
	//...
}

func (s *CreateUserUseCaseTestSuite) TestExecute_UserNotCreatedError() {
	//...
}

func TestCreateUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserUseCaseTestSuite))
}
