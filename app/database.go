package app

import (
	"MyGram/config"
	"MyGram/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func InitDB() *gorm.DB {
	envConfig := config.LoadConfig()

	dsn := "host=" + envConfig.DB_HOST + " user=" + envConfig.DB_USERNAME + " password=" + envConfig.DB_PASSWORD + " dbname=" + envConfig.DB_DATABASE + " port=" + strconv.Itoa(envConfig.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	helper.PanicIfError(err)

	return db
}
