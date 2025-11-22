package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ParseAndValidateBody(c *fiber.Ctx, target interface{}) error {

	if err := c.BodyParser(target); err != nil {
		return err
	}

	if err := GetValidator().Struct(target); err != nil {
		return err
	}

	return nil

}
