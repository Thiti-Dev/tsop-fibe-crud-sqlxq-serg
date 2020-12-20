package utils

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber"
)

// DoesReqBodyCanParsed is for checking if the Passed body can be parsed
func DoesReqBodyCanParsed(model interface{}, c *fiber.Ctx) error {
	req := model
	if err := c.BodyParser(req); err != nil {
		fmt.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
		return errors.New("Invalid JSON format")
	}
	return nil
}

// MakeErrorRespondIfAnyError does what it's called
func MakeErrorRespondIfAnyError(e error, c *fiber.Ctx){
	if e != nil{
		c.Status(400).JSON(&fiber.Map{
			"status":  400,
			"message": e.Error(),
		})
	}
}