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

	clients := app.Group("/clients")
	clients.Get("/", clientsHandler.main)
	clients.Get("/read", clientsHandler.read)
	clients.Post("/create", clientsHandler.create)
	clients.Put("/update", clientsHandler.update)
	clients.Delete("/delete", clientsHandler.delete)

	clientTypes := app.Group("/client-types")
	clientTypes.Get("/", clientTypesHandler.main)
	clientTypes.Get("/read", clientTypesHandler.read)
	clientTypes.Post("/create", clientTypesHandler.create)
	clientTypes.Put("/update", clientTypesHandler.update)
	clientTypes.Delete("/delete", clientTypesHandler.delete)

	app.Get("/", indexHandler.get)
	app.Get("/products", productsHandler.get)
	app.Get("/stock", stockHandler.get)
	app.Get("/supplers", supplersHandler.get)
	app.Get("/transactions", transactionsHandler.get)

	return app
}
