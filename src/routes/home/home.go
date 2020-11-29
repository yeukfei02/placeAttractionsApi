package home

import (
	"github.com/gofiber/fiber/v2"
	homeController "github.com/yeukfei02/placeAttractionsApi/src/controller/home"
)

// Routes func
func Routes(app *fiber.App) {
	app.Get("/", homeController.GetHome)
}
