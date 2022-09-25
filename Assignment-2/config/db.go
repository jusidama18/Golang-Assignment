package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)



func InitDB() *gorm.DB {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Println("config.env Not Found !")
	}

	dburl := os.Getenv("MYSQL_DB")

	db, err := gorm.Open("mysql", dburl)
	if err != nil {
		log.Panicf("DB ERROR : %s\n", err.Error())
	}
	return db
}