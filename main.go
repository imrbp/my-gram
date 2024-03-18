package main

import (
	"MyGram/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()
	app.Mount("", route.UserHandler())
	app.Mount("", route.PhotoHandler())
	app.Mount("", route.CommentHandler())
	app.Mount("", route.SocialMediaHandler())

	log.Fatal(app.Listen(":8080"))
}
