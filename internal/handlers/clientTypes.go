package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yvv4git/erp-fglaw/internal/config"
	"github.com/yvv4git/erp-fglaw/internal/forms"
	"gorm.io/gorm"
)

// ClientTypesHandler is handler.
type ClientTypesHandler struct {
	config config.Config
	db     *gorm.DB
}

func (h *ClientTypesHandler) main(c *fiber.Ctx) error {
	return c.Render("client_types/index", fiber.Map{
		"Title": "Main page",
	})
}

func (h *ClientTypesHandler) read(c *fiber.Ctx) error {
	form := new(forms.ClientTypes)

	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := forms.Validate(form); err != nil {
		return err
	}

	result, err := form.ReadPage(h.db)
	if err != nil {
		return err
	}

	return c.
		Status(200).
		JSON(result)
}

func (h *ClientTypesHandler) create(c *fiber.Ctx) error {
	form := new(forms.ClientTypes)

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

func (h *ClientTypesHandler) update(c *fiber.Ctx) error {
	form := new(forms.ClientTypes)

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

func (h *ClientTypesHandler) delete(c *fiber.Ctx) error {
	form := new(forms.ClientTypes)

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

// NewClientTypesHandler is used as constructor.
func NewClientTypesHandler(config config.Config, db *gorm.DB) *ClientTypesHandler {
	return &ClientTypesHandler{
		config: config,
		db:     db,
	}
}
