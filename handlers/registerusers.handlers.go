package handlers

import (
	"net/http"

	"github.com/Shoetan/db"
	"github.com/Shoetan/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


// Function to hash the password using bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
			return "", err
	}
	return string(hashedPassword), nil
}

func RegisterUser(ctx *gin.Context){
	// initialize db 
	db:=db.InitializeDb()

	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register user",
			"error": err,
		})

		return
	}

	// hash the password 
	hashedPassword, err:= hashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to register user",
				"error":   err.Error(),
		})
		return
}


	user.Password = hashedPassword 

	

	//create a new user in the database using GORM

	if result:= db.Create(&user); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register user",
			"error": result.Error.Error(),
		})
		return
	}

	//User registration successful

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"user": user,
	})
}

