package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/Shoetan/db"
	"github.com/Shoetan/models"
	"github.com/Shoetan/types"
	"github.com/gin-gonic/gin"
	"github.com/Shoetan/utils"
	
)



func LoginUser(ctx *gin.Context) {

	//initialize database
	db:= db.InitializeDb()

	// read request body from user 
	requestBody, err := ctx.GetRawData()

	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message":"Invalid request body",
			"status":err,
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

	

	compareResult := utils.ComparePassword(user.Password, payload.Password); 

	if compareResult != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message":"Invalid password",
			"status": compareResult.Error(),
		})
	} else {
		
		token,tokenErr:= utils.CreateToken(payload.Email)
	
		if tokenErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H {
				"message":"Invalid token",
				"token":token,
				"err":tokenErr.Error(),
			})
		}else{
			ctx.JSON(http.StatusOK, gin.H {
				"message":"Login successful",
				"token": token,
			})

		}	
	}	

}