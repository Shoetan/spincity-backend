package routers

import (
	"github.com/Shoetan/handlers"
	"github.com/gin-gonic/gin"
)

func CreateGenre ( router *gin.Engine) {
	router.POST("/apiV1/createGenre", handlers.CreateGenre)
}