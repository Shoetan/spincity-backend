package handlers

import (
  "encoding/json"
  "net/http"
  "github.com/Shoetan/db"
  "github.com/Shoetan/models"
  "github.com/gin-gonic/gin"
)

//Create the  handler function for this 

func CreateGenre(ctx *gin.Context) {
  // Initialize DB
  db := db.InitializeDb()

  // Get request body from the user
  genreBody, err := ctx.GetRawData()
	
  if err != nil {
    ctx.JSON(http.StatusBadGateway, gin.H{
      "message": "Failed to create genre",
      "error": err.Error(),
    })
    return
  }

  // Check the length of the request body is empty
  if len(genreBody) == 0 {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message": "Request body is empty",
    })
    return
  }

  // Make an instance of the genre model
  var genre models.Genre

  // Unmarshall the request body
  if err := json.Unmarshal(genreBody, &genre); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message": "Invalid request body or JSON format",
      "error": err,
    })
    return
  }

  // Check if a genre with the same title already exists

  if result := db.First(&genre, "title = ?", genre.Title); result.RowsAffected > 0 {
    ctx.JSON(http.StatusConflict, gin.H{
      "message": "Genre with the same title already exists",
      "error": "Duplicate genre title",
    })
    return
  }

  // Create the genre if it doesn't exist
  if result := db.Create(&genre); result.Error != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{
      "message": "Failed to create genre",
      "error": result.Error.Error(),
    })
  } else {
    ctx.JSON(http.StatusCreated, gin.H{
      "message": "Genre created",
      "genre": genre,
    })
  }
}