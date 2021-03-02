package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/yvv4git/erp-fglaw/internal/config"
)

// SetupWebServer is used for setup web server routes.
func SetupWebServer(config config.Config) *fiber.App {

	engine := django.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	indexHandler := NewIndexHandler(config)
	clientsHandler := NewClientsHandler(config)
	clientTypesHandler := NewClientTypesHandler(config)
	productsHandler := NewProductsHandler(config)
	stockHandler := NewStockHandler(config)
	supplersHandler := NewSupplersHandler(config)
	transactionsHandler := NewTransactionsHandler(config)

	app.Get("/", indexHandler.get)
	app.Get("/clients", clientsHandler.get)
	app.Get("/client-types", clientTypesHandler.get)
	app.Get("/products", productsHandler.get)
	app.Get("/stock", stockHandler.get)
	app.Get("/supplers", supplersHandler.get)
	app.Get("/transactions", transactionsHandler.get)

	return app
}
