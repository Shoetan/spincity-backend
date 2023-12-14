package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Shoetan/handlers"
)

func CreateGenre ( router *gin.Engine) {
	router.POST("/apiV1/createGenre", handlers.CreateGenre)
}