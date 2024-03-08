package route

import (
	"github.com/gofiber/fiber/v2"
)

func SocialMediaHandler() *fiber.App {
	app := fiber.New()

	socialMediaH := app.Group("/socialmedias")
	socialMediaH.Post("")
	socialMediaH.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello This is social media Handler")
	})
	socialMediaH.Put("/:socialMediaId")
	socialMediaH.Delete("/:socialMediaId")
	return app
}
