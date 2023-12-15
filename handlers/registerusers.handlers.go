package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Shoetan/db"
	"github.com/Shoetan/models"
	"github.com/Shoetan/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/Shoetan/types"
)


func RegisterUser(ctx *gin.Context){
	// initialize db 
	db:=db.InitializeDb()

	// Read the request body from user 
	requestBody, err := ctx.GetRawData()


	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message":"Failed to register user",
			"error": err.Error(),
		})
		return
	}
	//check the length of the request body is empty

	if len(requestBody) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Request body is empty",
		})
	}


	//make an instance of the user model 
	var user models.User

	if err := json.Unmarshal(requestBody, &user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format",
			"error": err,
		})

		return
	}

	// hash the password 
	hashedPassword, err:= utils.HashPassword(user.Password)

	hashedConfirmPassword, err:= utils.HashPassword(user.ConfrimPassword)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to register user",
				"error":   err.Error(),
		})
		return
}

//check db if user email exists
if result := db.First(&user, "email = ?", user.Email); result.RowsAffected > 0 {
	ctx.JSON(http.StatusConflict, gin.H {
		"message": "This email address already exists",
    "error": "Duplicate email address",
	})
}

	//updating user object wwith hashed password 
	user.Password = hashedPassword 
	user.ConfrimPassword = hashedConfirmPassword

	//create a new user in the database using GORM

	if result:= db.Create(&user); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register user",
			"error": result.Error.Error(),
		})
		return
	}

	responseUser := types.ResponseUser{
		Email:user.Email,
		Name: user.Name,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05 -070000"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05 -070000"),
		ID: strconv.FormatUint(uint64(user.ID), 10),
	}

	//User registration successful

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"user": responseUser,
	})
}

