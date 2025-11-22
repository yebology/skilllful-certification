package output

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/skillful-certification/app/dto/response"
)

func GetOutput(c *fiber.Ctx, status string, fiberStatus int, message string, data interface{}) error {

	response := response.ApiResponse{
		Status:  status,
		Message: message,
	}

	if data != nil {
		response.Data = data
	}

	return c.Status(fiberStatus).JSON(response)

}
