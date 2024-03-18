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

func CommentHandler() *fiber.App {

	app := fiber.New()

	db := app2.InitDB()
	validate := validator.New()
	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	controllerComment := controller.NewCommentController(validate, commentService)

	authMiddleware := keyauth.New(keyauth.Config{
		Validator: helper.AuthBearerValidation,
	})

	commentH := app.Group("/comments")
	commentH.Post("", authMiddleware, controllerComment.Create)
	commentH.Get("", authMiddleware, controllerComment.GetAll)
	commentH.Put("/:commentId", authMiddleware, controllerComment.Update)
	commentH.Delete("/:commentId", authMiddleware, controllerComment.Delete)
	return app
}
