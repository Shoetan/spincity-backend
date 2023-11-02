package db

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"

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

	return db
}