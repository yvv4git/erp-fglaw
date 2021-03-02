package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
)

// TransactionsHandler is handler.
type TransactionsHandler struct {
	config config.Config
}

func (h *TransactionsHandler) get(c *fiber.Ctx) error {
	return c.Render("transactions/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewTransactionsHandler is used as constructor.
func NewTransactionsHandler(config config.Config) *TransactionsHandler {
	return &TransactionsHandler{
		config: config,
	}
}
