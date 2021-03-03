package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"gorm.io/gorm"
)

// ProductsHandler is handler.
type ProductsHandler struct {
	config config.Config
	db     *gorm.DB
}

func (h *ProductsHandler) get(c *fiber.Ctx) error {
	return c.Render("products/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewProductsHandler is used as constructor.
func NewProductsHandler(config config.Config, db *gorm.DB) *ProductsHandler {
	return &ProductsHandler{
		config: config,
		db:     db,
	}
}
