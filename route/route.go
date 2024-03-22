package route

import (
	"MyGram/controller"
	"MyGram/helper"
	"MyGram/repository"
	"MyGram/service"
	"MyGram/storage"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func Route() *fiber.App {

	app := fiber.New()

	db := storage.InitDB()
	validate := validator.New()
	NewValidate := helper.NewValidator(validate)

	repositoryUser := repository.NewUserRepository(db)
	serviceUser := service.NewUserService(repositoryUser)
	controllerUser := controller.NewUserController(serviceUser, &NewValidate)

	authMiddleware := keyauth.New(keyauth.Config{
		Validator: helper.AuthBearerValidation,
	})

	userH := app.Group("/users")
	userH.Post("/register", controllerUser.Register)
	userH.Post("/login", controllerUser.Login)
	userH.Put("", authMiddleware, controllerUser.Update)
	userH.Delete("", authMiddleware, controllerUser.Delete)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	controllerPhoto := controller.NewPhotoController(photoService, &NewValidate)

	photoH := app.Group("/photos")
	photoH.Post("", authMiddleware, controllerPhoto.Post)
	photoH.Get("", authMiddleware, controllerPhoto.GetAll)
	photoH.Get("/:photoId", authMiddleware, controllerPhoto.Get)
	photoH.Put("/:photoId", authMiddleware, controllerPhoto.Update)
	photoH.Delete("/:photoId", authMiddleware, controllerPhoto.Delete)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	controllerComment := controller.NewCommentController(&NewValidate, commentService)

	commentH := app.Group("/comments")
	commentH.Post("", authMiddleware, controllerComment.Create)
	commentH.Get("", authMiddleware, controllerComment.GetAll)
	commentH.Get("/:commentId", authMiddleware, controllerComment.Get)
	commentH.Put("/:commentId", authMiddleware, controllerComment.Update)
	commentH.Delete("/:commentId", authMiddleware, controllerComment.Delete)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controller.NewSocialMediaController(&NewValidate, socialMediaService)

	socialMediaH := app.Group("/socialmedias")
	socialMediaH.Post("", authMiddleware, socialMediaController.Create)
	socialMediaH.Get("", authMiddleware, socialMediaController.GetAll)
	socialMediaH.Put("/:socialMediaId", authMiddleware, socialMediaController.Update)
	socialMediaH.Delete("/:socialMediaId", authMiddleware, socialMediaController.Delete)
	return app
}
