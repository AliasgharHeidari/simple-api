package handler

import (
	"github.com/AliasgharHeidari/simple-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func GopherMembers(c *fiber.Ctx) error {
	return c.JSON(model.Map_members)
}

func GopherTeacher(c *fiber.Ctx) error {
	return c.JSON(model.Map_teacher)
}

func GophersContact(c *fiber.Ctx) error {
	return c.JSON(model.Map_contact)
}

func GophersWebsite(c *fiber.Ctx) error {
	return c.JSON(model.Map_website)
}
