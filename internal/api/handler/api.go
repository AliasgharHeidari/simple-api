package handler

import (
	"github.com/gofiber/fiber/v2"
)

type members struct {
	Names string
	Age   string
}

var map_members = map[int]members{
	1: {Names: "Aliasghar", Age: "18"},
	2: {Names: "Shahrzad", Age: "17"},
	3: {Names: "AmirHosein", Age: "21"},
}

func GopherMembers(c *fiber.Ctx) error {
	c.SendString(map_members[1].Names)
	return c.SendString(map_members[1].Age)
}

func GopherTeacher(c *fiber.Ctx) error {
	return c.SendString("Mahmoud")
}
