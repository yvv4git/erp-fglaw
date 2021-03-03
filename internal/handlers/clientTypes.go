package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"gorm.io/gorm"
)

// ClientTypesHandler is handler.
type ClientTypesHandler struct {
	config config.Config
	db     *gorm.DB
}

func (h *ClientTypesHandler) get(c *fiber.Ctx) error {
	return c.Render("clienttypes/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewClientTypesHandler is used as constructor.
func NewClientTypesHandler(config config.Config, db *gorm.DB) *ClientTypesHandler {
	return &ClientTypesHandler{
		config: config,
		db:     db,
	}
}
