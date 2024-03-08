package main

import (
	"MyGram/route"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	app.Mount("", route.UserHandler())
	//app.Mount("", route.PhotoHandler())
	//app.Mount("", route.CommentHandler())
	//app.Mount("", route.SocialMediaHandler())
	log.Fatal(app.Listen(":8080"))
	//Db := app.InitDB()
	//err := Db.AutoMigrate(entity.User{}, entity.SocialMedia{}, entity.Comment{}, entity.Photo{})
	//
	//user := entity.UserCreateRequest{
	//	Username: "imrbp",
	//	Email:    "imrbp@gmail.com",
	//	Password: "flsdjflkdsjf",
	//	Age:      10,
	//}

	//user := entity.User{
	//	Id: 5,
	//}

	//repo := repository.NewUserRepository(Db)
	//println(result.Username )
	//result := repository.NewUserRepository(Db).Update(context.Background(), user)
	//result := repository.NewUserRepository(Db).FindMatch(context.Background(), user)
	//result := repository.NewUserRepository(Db).Delete(context.Background(), user)

	//validation := new(validator.Validate)
	//ser := service.NewUserService(*repo)
	//cont := controller.NewUserController(*ser, validation)

	//helper.PanicIfError(err)
	//println(result.Id)

}
