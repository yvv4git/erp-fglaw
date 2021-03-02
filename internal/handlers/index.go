package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
)

// IndexHandler is handler.
type IndexHandler struct {
	config config.Config
}

func (h *IndexHandler) get(c *fiber.Ctx) error {
	return c.Render("index/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewIndexHandler is used as constructor.
func NewIndexHandler(config config.Config) *IndexHandler {
	return &IndexHandler{
		config: config,
	}
}
