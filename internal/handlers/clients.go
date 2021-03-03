package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"gorm.io/gorm"
)

// ClientsHandler is handler.
type ClientsHandler struct {
	config config.Config
	db     *gorm.DB
}

func (h *ClientsHandler) get(c *fiber.Ctx) error {
	return c.Render("clients/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewClientsHandler is used as constructor.
func NewClientsHandler(config config.Config, db *gorm.DB) *ClientsHandler {
	return &ClientsHandler{
		config: config,
		db:     db,
	}
}
