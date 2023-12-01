package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/Shoetan/handlers"
	"github.com/gin-gonic/gin"
)

func ReturnGenre (router *gin.Engine) {
	// Add cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:[]string{ "http://localhost:3000"     },
		AllowMethods: []string{"GET"},
    AllowHeaders: []string{"Content-Type, Authorization"},
	}))
	router.GET("/apiV1/returngenre", handlers.ReturnGenreFromDb)
}