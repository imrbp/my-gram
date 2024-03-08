package route

import (
	"github.com/gofiber/fiber/v2"
)

func PhotoHandler() *fiber.App {
	app := fiber.New()

	photoH := app.Group("/photos")
	photoH.Post("")
	photoH.Get("")
	photoH.Put("/:photoId")
	photoH.Delete("/:photoId")
	return app
}
