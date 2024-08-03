package response

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

type ErrorValidateResponse struct {
	Success bool                `json:"success"`
	Status  int                 `json:"status"`
	Error   []map[string]string `json:"error"`
}

func BadRequestErrors(c *fiber.Ctx, errors []map[string]string) {
	c.Status(fiber.StatusBadRequest).JSON(ErrorValidateResponse{
		Success: false,
		Status:  400,
		Error:   errors,
	})
}

func BadRequest(c *fiber.Ctx, err error) {
	c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
		Success: false,
		Status:  400,
		Error:   err.Error(),
	})
}

func InternalServerError(c *fiber.Ctx, err error) {
	c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
		Success: false,
		Status:  500,
		Error:   err.Error(),
	})
}
