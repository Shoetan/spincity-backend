package handlers

import (
  //"encoding/json"
  "net/http"
  "github.com/Shoetan/db"
  "github.com/Shoetan/models"
  "github.com/gin-gonic/gin"
)

// create handler function to return all genre from the database

func ReturnGenreFromDb (ctx *gin.Context) {

	//initial db 

	db:= db.InitializeDb()

	var genres  []models.Genre

	if result := db.Find(&genres); result.Error != nil {
    ctx.JSON(http.StatusNotFound, gin.H{
      "message": "No records found",
      "error": result.Error.Error(),
    })
  } else {
    ctx.JSON(http.StatusOK, gin.H{
      "message": "Fetched genres",
      "genre": genres,
    })
  }
}