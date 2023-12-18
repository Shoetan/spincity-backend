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

var secretKey = []byte(getSecretKey("SECRET_KEY"))

func CreateToken(email string) (string, error) {
	
	
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

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
