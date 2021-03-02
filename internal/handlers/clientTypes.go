package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
)

// ClientTypesHandler is handler.
type ClientTypesHandler struct {
	config config.Config
}

func (h *ClientTypesHandler) get(c *fiber.Ctx) error {
	return c.Render("clienttypes/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewClientTypesHandler is used as constructor.
func NewClientTypesHandler(config config.Config) *ClientTypesHandler {
	return &ClientTypesHandler{
		config: config,
	}
}
