package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
)

func getSecretKey(key string) string {
	err:= godotenv.Load(".env")

	if err != nil {
		log.Fatalln(err)
	}
	return os.Getenv(key)
}

func CreateToken(email string) (string, error) {
	secretKey:= []byte(getSecretKey("SECRET_KEY"))
	fmt.Println(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{ 
		"userEmail": email, 
		"exp": time.Now().Add(time.Hour * 72).Unix(), 
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
	return "", err
	}

return tokenString, nil
}