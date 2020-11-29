package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/yeukfei02/placeAttractionsApi/src/routes/home"
	"github.com/yeukfei02/placeAttractionsApi/src/routes/place"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(compress.New())

	home.Routes(app)
	place.Routes(app)

	app.Listen(":3000")
}
