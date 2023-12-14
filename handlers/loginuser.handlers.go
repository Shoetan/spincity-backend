package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Shoetan/db"
	"github.com/Shoetan/models"
	"github.com/Shoetan/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Function to  compare hashed passwords against its plain text equivalent

func comparePassword(hashedPassword, plainPassword string) (error){
	return  bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func LoginUser(ctx *gin.Context) {

	//initialize database
	db:= db.InitializeDb()

	// read request body from user 
	requestBody, err := ctx.GetRawData()

	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message":"Invalid request body",
			"status":"fail",
		})
	}

	//make an instance of the request body

	var payload types.LoginPayload

	var user models.User

	json.Unmarshal(requestBody, &payload)

	// get hashed password from db

	if result := db.First(&user, "email = ?" , payload.Email); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message":"Invalid email address or password",
			"status":"fail",
		})
	}

	

	err = comparePassword(user.Password, payload.Password); 

	if err !=nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message":"Invalid password",
			"status":"failure",
		})
	}


}