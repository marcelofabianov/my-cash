package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/marcelofabianov/my-cash/internal/adapter/http/request"
	"github.com/marcelofabianov/my-cash/internal/adapter/http/response"
	dError "github.com/marcelofabianov/my-cash/internal/domain/error"
	"github.com/marcelofabianov/my-cash/internal/port/inbound"
)

func CreateUserHandler(c *fiber.Ctx) error {
	var data inbound.CreateUserRequest

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	isValid := request.Validate(c, data)
	if isValid {
		inputS := inbound.CreateUserServiceInput{
			CreateUserUseCaseInput: inbound.CreateUserUseCaseInput{
				Document: data.Document,
				Name:     data.Name,
				Email:    data.Email,
				Password: data.Password,
			},
		}

		service := c.Locals("userService").(inbound.UserService)

		outputS, err := service.CreateUser(c.Context(), inputS)
		if err != nil {
			if dError.IsUserExistsError(err) {
				response.BadRequest(c, err)
				return nil
			}

			response.InternalServerError(c, err)
			return nil
		}

		presenter := inbound.UserPresenter{
			ID:        outputS.User.ID.String(),
			Document:  outputS.User.Document.String(),
			Name:      outputS.User.Name,
			Email:     outputS.User.Email.String(),
			Enabled:   outputS.User.Enabled.Bool(),
			CreatedAt: outputS.User.CreatedAt.String(),
			UpdatedAt: outputS.User.UpdatedAt.String(),
		}

		response.Created(c, presenter)
	}

	return nil
}
