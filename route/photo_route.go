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

func PhotoHandler() *fiber.App {
	app := fiber.New()

	db := app2.InitDB()
	validate := validator.New()
	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	controllerPhoto := controller.NewPhotoController(photoService, validate)

	authMiddleware := keyauth.New(keyauth.Config{
		Validator: helper.AuthBearerValidation,
	})

	photoH := app.Group("/photos")
	photoH.Post("", authMiddleware, controllerPhoto.Post)
	photoH.Get("", authMiddleware, controllerPhoto.GetAll)
	photoH.Put("/:photoId", authMiddleware, controllerPhoto.Update)
	photoH.Delete("/:photoId", authMiddleware, controllerPhoto.Delete)
	return app
}
