package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"gorm.io/gorm"
)

// StockHandler is handler.
type StockHandler struct {
	config config.Config
	db     *gorm.DB
}

func (h *StockHandler) get(c *fiber.Ctx) error {
	return c.Render("stock/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewStockHandler is used as constructor.
func NewStockHandler(config config.Config, db *gorm.DB) *StockHandler {
	return &StockHandler{
		config: config,
		db:     db,
	}
}
