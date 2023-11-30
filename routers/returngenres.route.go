package routers

import (
	"github.com/Shoetan/handlers"
	"github.com/gin-gonic/gin"
)

func ReturnGenre (router *gin.Engine) {
	router.GET("/apiV1/returngenre", handlers.ReturnGenreFromDb)
}