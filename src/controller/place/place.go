package place

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeukfei02/placeAttractionsApi/src/common"
)

// CreatePlaceAttractionsRequestBody struct
type CreatePlaceAttractionsRequestBody struct {
	Name     string   `json:"name"`
	Rating   float64  `json:"rating"`
	Location Location `json:"location"`
}

// Location struct
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// CreatePlaceAttractions func
func CreatePlaceAttractions(c *fiber.Ctx) error {
	placeAttractionsRequestBody := new(CreatePlaceAttractionsRequestBody)

	err := c.BodyParser(placeAttractionsRequestBody)
	common.CheckErr(err)

	// name := placeAttractionsRequestBody.Name
	// rating := placeAttractionsRequestBody.Rating
	// lat := placeAttractionsRequestBody.Location.Lat
	// lng := placeAttractionsRequestBody.Location.Lng

	return c.Status(201).JSON(fiber.Map{
		"message": "Create Place Attractions",
	})
}

// GetPlaceAttractions func
func GetPlaceAttractions(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Get Place Attractions",
	})
}

// GetPlaceAttractionsByID func
func GetPlaceAttractionsByID(c *fiber.Ctx) error {
	var response error

	id := c.Params("id")

	if len(id) > 0 {
		response = c.Status(200).JSON(fiber.Map{
			"message": "Get Place Attractions By Id",
		})
	}

	return response
}

// UpdatePlaceAttractionsByID func
func UpdatePlaceAttractionsByID(c *fiber.Ctx) error {
	var response error

	id := c.Params("id")

	if len(id) > 0 {
		response = c.Status(200).JSON(fiber.Map{
			"message": "Update Place Attractions By Id",
		})
	}

	return response
}

// DeletePlaceAttractionsByID func
func DeletePlaceAttractionsByID(c *fiber.Ctx) error {
	var response error

	id := c.Params("id")

	if len(id) > 0 {
		response = c.Status(200).JSON(fiber.Map{
			"message": "Delete Place Attractions By Id",
		})
	}

	return response
}
