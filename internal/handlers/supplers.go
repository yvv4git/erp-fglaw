package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
)

//SupplersHandler is handler.
type SupplersHandler struct {
	config config.Config
}

func (h *SupplersHandler) get(c *fiber.Ctx) error {
	return c.Render("supplers/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewSupplersHandler is used as constructor.
func NewSupplersHandler(config config.Config) *SupplersHandler {
	return &SupplersHandler{
		config: config,
	}
}
