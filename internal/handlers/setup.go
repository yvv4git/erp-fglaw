package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"gorm.io/gorm"
)

// SetupWebServer is used for setup web server routes.
func SetupWebServer(config config.Config, db *gorm.DB) *fiber.App {

	engine := django.New("./web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	indexHandler := NewIndexHandler(config, db)
	clientsHandler := NewClientsHandler(config, db)
	clientTypesHandler := NewClientTypesHandler(config, db)
	productsHandler := NewProductsHandler(config, db)
	stockHandler := NewStockHandler(config, db)
	supplersHandler := NewSupplersHandler(config, db)
	transactionsHandler := NewTransactionsHandler(config, db)

	app.Get("/", indexHandler.get)
	app.Get("/clients", clientsHandler.get)
	app.Get("/client-types", clientTypesHandler.get)
	app.Get("/products", productsHandler.get)
	app.Get("/stock", stockHandler.get)
	app.Get("/supplers", supplersHandler.get)
	app.Get("/transactions", transactionsHandler.get)

	return app
}
