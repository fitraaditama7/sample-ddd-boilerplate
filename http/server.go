package http

import (
	"ddd-boilerplate/config"
	"ddd-boilerplate/domain/sample"
	"ddd-boilerplate/http/handler"
	"ddd-boilerplate/http/router"
	"ddd-boilerplate/pkg/fiber"
	"ddd-boilerplate/pkg/metrics"
	sampleOutboundRepo "ddd-boilerplate/repositories/outbound/sample"
	"github.com/gofiber/swagger"
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

// @title New Simobi+ API
// @version 2.0
// @description This is a API docs for new simobi+.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https
func StartServer(config *config.Config) {
	outboundMetrics := metrics.SetupOutboundMetric()
	prometheus.MustRegister(outboundMetrics)

	sampleOutbound := sampleOutboundRepo.NewOutbound(outboundMetrics)

	sampleService := sample.NewSampleService(sampleOutbound)
	sampleHandler := handler.NewHandler(sampleService)

	app := fiber.InitFiberApp()
	app.Get("/swagger/*", swagger.HandlerDefault)

	router.SetupRoutes(app, *sampleHandler)
	log.Fatal(app.Listen(config.App.Port))
}
