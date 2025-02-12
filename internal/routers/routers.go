package routers

import (
	"github.com/alwialdi9/be-jajanskuy/internal/controllers"
	"github.com/alwialdi9/be-jajanskuy/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Route() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(middlewares.LoggerMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	api.Post("/init", controllers.InitController)

	return app
}
