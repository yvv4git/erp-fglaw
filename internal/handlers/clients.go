package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"github.com/yvv4git/erp-fglaw/internal/forms"
	"gorm.io/gorm"
)

// ClientsHandler is handler.
type ClientsHandler struct {
	config config.Config
	db     *gorm.DB
}

func (h *ClientsHandler) main(c *fiber.Ctx) error {
	return c.Render("clients/index", fiber.Map{
		"Title": "Main page",
	})
}

func (h *ClientsHandler) read(c *fiber.Ctx) error {
	form := new(forms.Clients)

	if err := c.QueryParser(form); err != nil {
		return err
	}

	if err := forms.Validate(form); err != nil {
		return err
	}

	clients, err := form.ReadPage(h.db)
	if err != nil {
		return err
	}

	count, err := form.Count(h.db)
	if err != nil {
		return err
	}

	return c.
		Status(200).
		JSON(fiber.Map{
			"clients": clients,
			"count":   count,
		})
}

func (h *ClientsHandler) create(c *fiber.Ctx) error {
	form := new(forms.Clients)

	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := forms.Validate(form); err != nil {
		return err
	}

	if err := form.Create(h.db); err != nil {
		return err
	}

	return c.
		Status(200).
		JSON(&fiber.Map{
			"success": true,
		})
}

func (h *ClientsHandler) update(c *fiber.Ctx) error {
	form := new(forms.Clients)

	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := forms.Validate(form); err != nil {
		return err
	}

	if err := form.Update(h.db); err != nil {
		return err
	}

	return c.
		Status(200).
		JSON(&fiber.Map{
			"success": true,
		})
}

func (h *ClientsHandler) delete(c *fiber.Ctx) error {
	form := new(forms.Clients)

	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := forms.Validate(form); err != nil {
		return err
	}

	if err := form.Delete(h.db); err != nil {
		return err
	}

	return c.
		Status(200).
		JSON(&fiber.Map{
			"success": true,
		})
}

// NewClientsHandler is used as constructor.
func NewClientsHandler(config config.Config, db *gorm.DB) *ClientsHandler {
	return &ClientsHandler{
		config: config,
		db:     db,
	}
}
