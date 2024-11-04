package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamesmpuya/tapin-go-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
