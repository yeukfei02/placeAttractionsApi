package home

import "github.com/gofiber/fiber/v2"

// GetHome func
func GetHome(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "placeAttractionsApi",
	})
}
