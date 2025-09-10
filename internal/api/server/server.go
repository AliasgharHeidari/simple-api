package apiserver

import (
	"os"
	"github.com/AliasgharHeidari/simple-api/internal/api/handler"
	"github.com/gofiber/fiber/v2"
)

func Start() {
app := fiber.New()

	app.Get("/gophers/members",handler.GopherMembers)
	app.Get("/gophers/teacher",handler.GopherTeacher)
	app.Get("/gophers/contact",handler.GophersContact)
	app.Get("/gophers/website",handler.GophersWebsite)

	port :=os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Listen(":"+port)
}