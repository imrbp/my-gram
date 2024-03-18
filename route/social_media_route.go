package route

import (
	app2 "MyGram/app"
	"MyGram/controller"
	"MyGram/helper"
	"MyGram/repository"
	"MyGram/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func SocialMediaHandler() *fiber.App {
	app := fiber.New()

	db := app2.InitDB()
	validate := validator.New()
	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controller.NewSocialMediaController(validate, socialMediaService)

	authMiddleware := keyauth.New(keyauth.Config{
		Validator: helper.AuthBearerValidation,
	})

	socialMediaH := app.Group("/socialmedias")
	socialMediaH.Post("", socialMediaController.Create)
	socialMediaH.Get("", authMiddleware, socialMediaController.GetAll)
	socialMediaH.Put("/:socialMediaId", socialMediaController.Update)
	socialMediaH.Delete("/:socialMediaId", socialMediaController.Delete)
	return app
}
