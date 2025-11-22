package utils

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ConvertToNum(c *fiber.Ctx, idStr string) (int, error) {

	id := c.Params(idStr)

	idNum, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return idNum, nil

}

func ConvertStrToDate(date string) (time.Time, error) {

	converted, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}

	return converted, nil

}

func ConvertDateToStr(date time.Time) string {

	return date.Format("2006-01-02")

}
