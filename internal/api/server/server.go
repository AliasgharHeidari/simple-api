package apiserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AliasgharHeidari/simple-api/internal/api/handler"
)

func Start() {
app := fiber.New()

	app.Get("/gophers/members", handler.GopherMembers)
	app.Get("/gophers/teacher",handler.GopherTeacher)

	app.Listen(":8080")
}