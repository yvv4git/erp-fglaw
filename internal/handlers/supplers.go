package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"gorm.io/gorm"
)

//SupplersHandler is handler.
type SupplersHandler struct {
	config config.Config
	db     *gorm.DB
}

func (h *SupplersHandler) get(c *fiber.Ctx) error {
	return c.Render("supplers/index", fiber.Map{
		"Title": "Main page",
	})
}

// NewSupplersHandler is used as constructor.
func NewSupplersHandler(config config.Config, db *gorm.DB) *SupplersHandler {
	return &SupplersHandler{
		config: config,
		db:     db,
	}
}
