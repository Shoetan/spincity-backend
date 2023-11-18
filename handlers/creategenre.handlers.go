package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/Shoetan/db"
	"github.com/Shoetan/models"
	"github.com/gin-gonic/gin"
)

//Createa the  handler function for this 

func CreateGenre ( ctx *gin.Context) {
	//initialize db 

	db :=db.InitializeDb()

	//get request body from the user

	genreBody , err := ctx.GetRawData()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H {
			"message": "Failed to create genre",
			"error": err.Error(),
		})
		return
	}

	//check the length of the request body is empty

	if len(genreBody) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Request body is empty",
		})
	}

	//make an instance of the genre model 

	var genre models.Genre


	//unmarshall the request body

	if err := json.Unmarshal(genreBody, &genre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H {
			"message": "Invalid request body or Json format",
			"error": err,
		})

		return
	}

	if result := db.Create(&genre); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"message": "Failed to create Genre",
			"error": result.Error.Error(),
		} )
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Genre created",
			"genre": genre,
		})
	}

}


