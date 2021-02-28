package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/yvv4git/erp-fglaw/internal/config"
)

var cfg *config.Config

// SetupWebServer is used for setup web server routes.
func SetupWebServer(config *config.Config) *fiber.App {
	cfg = config

	engine := django.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", indexGet)

	return app
}
