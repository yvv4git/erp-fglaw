package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
)

// StockHandler is handler.
type StockHandler struct {
	config config.Config
}

func (h *StockHandler) get(c *fiber.Ctx) error {
	return c.Render("stock/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewStockHandler is used as constructor.
func NewStockHandler(config config.Config) *StockHandler {
	return &StockHandler{
		config: config,
	}
}
