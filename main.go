package main

import (
	"log"

	"api-gateway/config"
	"api-gateway/handlers"
	"api-gateway/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()

	// Public endpoint
	app.Post("/auth/guest",
		limiter.New(middleware.RateLimit()),
		handlers.GuestToken,
	)

	// Protected API
	api := app.Group("/api",
		middleware.JWTProtected(),
		limiter.New(middleware.RateLimit()),
	)

	api.All("/content/*",
		handlers.ProxyTo(config.BackendServices["content"]),
	)

	log.Fatal(app.Listen(":8080"))
}
