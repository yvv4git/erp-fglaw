package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func indexGet(c *fiber.Ctx) error {
	log.Println(cfg)
	return c.Render("index/index", fiber.Map{
		"Title": "Main page",
	})
}
