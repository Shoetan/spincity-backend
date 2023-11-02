package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Shoetan/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDbPassword (key string) string {
	err:= godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
	}
	return os.Getenv(key)
}

func InitializeDb() *gorm.DB{
	dbPassword:= getDbPassword("POSTGRES_PASSWORD")
	dbURL:= fmt.Sprintf("postgres://postgres:%s@localhost:5432/spincity", dbPassword)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Connected to database")
	}


// create a table in the db with the struct of VinylRecord
	db.AutoMigrate(&models.VinylRecord{})
	return db
}