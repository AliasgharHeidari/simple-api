package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AliasgharHeidari/simple-api/internal/model"
)



func GopherMembers(c *fiber.Ctx) error {
	return c.JSON(model.Map_members)
}


func GopherTeacher(c *fiber.Ctx) error {
	return c.JSON(model.Map_teacher)
}
