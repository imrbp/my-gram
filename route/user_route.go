package route

import (
	app2 "MyGram/app"
	controller2 "MyGram/controller"
	repository2 "MyGram/repository"
	"MyGram/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandler() *fiber.App {
	app := fiber.New()

	db := app2.InitDB()
	validate := validator.New()
	repositoryUser := repository2.NewUserRepository(db)
	serviceUser := service.NewUserService(repositoryUser)
	controllerUser := controller2.NewUserController(serviceUser, validate)
	userH := app.Group("/users")
	userH.Post("/register", controllerUser.Register)
	//userH.Post("/login")
	//userH.Put("/")
	return app
}
