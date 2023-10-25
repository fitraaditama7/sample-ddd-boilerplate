package router

import (
	"ddd-boilerplate/http/handler"
	"ddd-boilerplate/http/middleware"
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/module/apmfiber/v2"
)

func SetupRoutes(app *fiber.App, sampleHandler handler.Handler) {
	app.Use(apmfiber.Middleware())
	app.Use(middleware.Logger())
	sampleRouter := app.Group("/sample")
	sampleRouter.Get("/:id", sampleHandler.GetSampleByID)
}
