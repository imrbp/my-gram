package main

import (
	"MyGram/config"
	"MyGram/model/entity"
	"MyGram/route"
	"MyGram/storage"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	db := storage.InitDB()

	loadConfig := config.LoadConfig()

	if loadConfig.DB_AUTO_MIGRATE {
		err := db.AutoMigrate(entity.User{}, entity.Photo{}, entity.SocialMedias{}, entity.Comment{})
		if err != nil {
			panic(err)
		}
		//Default name gorm is social_media. so it's need to change
		err = db.Migrator().RenameTable(entity.SocialMedias{}, "social_medias")
		if err != nil {
			panic(err)
		}
		fmt.Println("Auto Migrate has been running. Application will close. please change it to false and run again")
	}

	app := fiber.New()
	app.Mount("", route.Route())
	log.Fatal(app.Listen(":8080"))
}
