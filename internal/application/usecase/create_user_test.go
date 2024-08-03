package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/marcelofabianov/my-cash/internal/application/usecase"
	inboundMock "github.com/marcelofabianov/my-cash/internal/port/inbound/mock"
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
	//...
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
