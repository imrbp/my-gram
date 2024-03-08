package route

import (
	"github.com/gofiber/fiber/v2"
)

func CommentHandler() *fiber.App {

	app := fiber.New()

	commentH := app.Group("/comments")
	commentH.Post("")
	commentH.Get("")
	commentH.Put("/:commentId")
	commentH.Delete("/:commentId")
	return app
}
