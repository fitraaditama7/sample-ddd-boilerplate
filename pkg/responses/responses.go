package responses

import (
	"ddd-boilerplate/pkg/custom_error"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type ResponseData struct {
	Code    string `json:"response_code"`
	Message string `json:"response_message"`
}

func Error(c *fiber.Ctx, err error) error {
	var re *custom_error.HTTPError
	if errors.As(err, &re) {
		return c.Status(re.StatusCode).JSON(ResponseData{
			Code:    re.ResponseCode,
			Message: re.Message,
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
		Code:    "Internal Server Error",
		Message: "Internal Server Error",
	})
}
