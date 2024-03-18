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

func UserHandler() *fiber.App {
	app := fiber.New()

	db := app2.InitDB()
	validate := validator.New()
	repositoryUser := repository.NewUserRepository(db)
	serviceUser := service.NewUserService(repositoryUser)
	controllerUser := controller.NewUserController(serviceUser, validate)

	authMiddleware := keyauth.New(keyauth.Config{
		Validator: helper.AuthBearerValidation,
	})

	userH := app.Group("/users")
	userH.Post("/register", controllerUser.Register)
	userH.Post("/login", controllerUser.Login)
	userH.Put("", authMiddleware, controllerUser.Update)
	userH.Delete("", authMiddleware, controllerUser.Delete)
	return app
}
