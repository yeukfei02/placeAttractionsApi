package place

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yeukfei02/placeAttractionsApi/src/common"
	placeAttractionsModel "github.com/yeukfei02/placeAttractionsApi/src/model/placeAttractions"
	placeServices "github.com/yeukfei02/placeAttractionsApi/src/services/place"
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
	var response error

	placeAttractionsRequestBody := new(CreatePlaceAttractionsRequestBody)

	err := c.BodyParser(placeAttractionsRequestBody)
	common.CheckErr(err)

	name := placeAttractionsRequestBody.Name
	rating := placeAttractionsRequestBody.Rating
	lat := placeAttractionsRequestBody.Location.Lat
	lng := placeAttractionsRequestBody.Location.Lng

	if len(name) > 0 && rating > 0 && lat != 0 && lng != 0 {
		placeServices.CreatePlaceAttractions(name, rating, lat, lng)
		response = c.Status(201).JSON(fiber.Map{
			"message": "Create Place Attractions",
		})
	} else {
		response = c.Status(400).JSON(fiber.Map{
			"message": "Create Place Attractions error, please enter name/rating/location",
		})
	}

	return response
}

// GetPlaceAttractions func
func GetPlaceAttractions(c *fiber.Ctx) error {
	var response error

	placeAttractionsList := placeServices.GetPlaceAttractions()

	if len(placeAttractionsList) > 0 {
		response = c.Status(200).JSON(fiber.Map{
			"message":          "Get Place Attractions",
			"placeAttractions": placeAttractionsList,
		})
	} else {
		placeAttractionsList := []placeAttractionsModel.PlaceAttractions{}
		response = c.Status(200).JSON(fiber.Map{
			"message":          "Get Place Attractions",
			"placeAttractions": placeAttractionsList,
		})
	}

	return response
}

// GetPlaceAttractionsByID func
func GetPlaceAttractionsByID(c *fiber.Ctx) error {
	var response error

	id := c.Params("id")
	if len(id) > 0 {
		placeAttractions := placeServices.GetPlaceAttractionsByID(id)
		if placeAttractions != nil {
			response = c.Status(200).JSON(fiber.Map{
				"message":         "Get Place Attractions By Id",
				"placeAttraction": placeAttractions,
			})
		}
	}

	return response
}

// UpdatePlaceAttractionsByID func
func UpdatePlaceAttractionsByID(c *fiber.Ctx) error {
	var response error

	id := c.Params("id")
	if len(id) > 0 {
		placeAttractionsRequestBody := new(CreatePlaceAttractionsRequestBody)

		err := c.BodyParser(placeAttractionsRequestBody)
		common.CheckErr(err)

		name := placeAttractionsRequestBody.Name
		rating := placeAttractionsRequestBody.Rating
		lat := placeAttractionsRequestBody.Location.Lat
		lng := placeAttractionsRequestBody.Location.Lng

		placeServices.UpdatePlaceAttractionsByID(id, name, rating, lat, lng)
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
		placeServices.DeletePlaceAttractionsByID(id)
		response = c.Status(200).JSON(fiber.Map{
			"message": "Delete Place Attractions By Id",
		})
	}

	return response
}
