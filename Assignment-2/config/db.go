package config

import (
	// "Assignment-2/repositories/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func createDB (db *gorm.DB) {
	db.Exec(`CREATE DATABASE IF NOT EXISTS orders_by`)
	db.Exec("USE orders_by")
}

func InitDB() *gorm.DB {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Println("config.env Not Found !")
	}

	dburl := os.Getenv("DB_URL")
	parsedb, err := pq.ParseURL(dburl)
	if err != nil {
		log.Panicf("DB ERROR : %s\n", err.Error())
	}
	db, err := gorm.Open(postgres.Open(parsedb), &gorm.Config{})
	if err != nil {
		log.Panicf("DB ERROR : %s\n", err.Error())
	}
	createDB(db)
	// db.AutoMigrate(models.Orders{}, models.Items{})
	return db
}