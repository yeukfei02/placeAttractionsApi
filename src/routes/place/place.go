package place

import (
	"github.com/gofiber/fiber/v2"
	placeController "github.com/yeukfei02/placeAttractionsApi/src/controller/place"
)

// Routes func
func Routes(app *fiber.App) {
	app.Post("/api/place-attractions", placeController.CreatePlaceAttractions)
	app.Get("/api/place-attractions", placeController.GetPlaceAttractions)
	app.Get("/api/place-attractions/:id", placeController.GetPlaceAttractionsByID)
	app.Put("/api/place-attractions/:id", placeController.UpdatePlaceAttractionsByID)
	app.Delete("/api/place-attractions/:id", placeController.DeletePlaceAttractionsByID)
}
