package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/yeukfei02/placeAttractionsApi/src/db"
	"github.com/yeukfei02/placeAttractionsApi/src/routes/home"
	"github.com/yeukfei02/placeAttractionsApi/src/routes/place"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(logger.New())

	db.ConnectDB()

	home.Routes(app)
	place.Routes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	app.Listen(port)
}
