package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"gorm.io/gorm"
)

// IndexHandler is handler.
type IndexHandler struct {
	config config.Config
	db     *gorm.DB
}

func (h *IndexHandler) get(c *fiber.Ctx) error {
	return c.Render("index/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewIndexHandler is used as constructor.
func NewIndexHandler(config config.Config, db *gorm.DB) *IndexHandler {
	return &IndexHandler{
		config: config,
		db:     db,
	}
}
