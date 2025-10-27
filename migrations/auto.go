package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/lowerKamaCase/product/pkg/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&product.Product{})
}
