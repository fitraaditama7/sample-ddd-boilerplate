package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitFiberApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	return app
}
