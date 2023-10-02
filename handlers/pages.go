package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func IndexPage(c *fiber.Ctx) error {
	if c.Locals("htmx") == true {
		return c.Render("index-content", nil)
	} else {
		return c.Render("index", nil)
	}
}
