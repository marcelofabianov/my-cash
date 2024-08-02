package container

import (
	"database/sql"

	"go.uber.org/dig"

	"github.com/marcelofabianov/my-cash/internal/adapter/postgres"
	"github.com/marcelofabianov/my-cash/internal/application/service"
	"github.com/marcelofabianov/my-cash/internal/application/usecase"
	"github.com/marcelofabianov/my-cash/internal/port/inbound"
	"github.com/marcelofabianov/my-cash/internal/port/outbound"
	"github.com/marcelofabianov/my-cash/pkg/hasher"
	"github.com/marcelofabianov/my-cash/pkg/logger"
)

type UserContainer struct {
	*dig.Container
}

func NewUserContainer(db *sql.DB, logger logger.Logger) *UserContainer {
	container := dig.New()

	userRegisterPackage(container)
	userRegisterRepositories(container, db)
	userRegisterUseCases(container)
	userRegisterServices(container, logger)

	return &UserContainer{container}
}

func userRegisterPackage(c *dig.Container) {
	c.Provide(func() inbound.PasswordHasher {
		return hasher.NewHasher()
	})
}

func userRegisterRepositories(c *dig.Container, db *sql.DB) {
	c.Provide(func() outbound.UserRepository {
		return postgres.NewUserRepository(db)
	})
}

func userRegisterUseCases(c *dig.Container) {
	c.Provide(func(r outbound.UserRepository, h inbound.PasswordHasher) inbound.CreateUserUseCase {
		return usecase.NewCreateUserUseCase(r, h)
	})
}

func userRegisterServices(c *dig.Container, l logger.Logger) {
	c.Provide(func(uc inbound.CreateUserUseCase) inbound.UserService {
		return service.NewUserService(l, uc)
	})
}
