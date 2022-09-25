package main

import (
	"Assignment-2/config"
	"Assignment-2/repositories"
)

func main() {
	db := config.InitDB()
	repositories.MakeOrderRepo(db)

}