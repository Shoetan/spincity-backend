package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/Shoetan/handlers"
	"github.com/gin-gonic/gin"
)

func LoginUser (router *gin.Engine) {

	// Add cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:[]string{ "http://localhost:9999"     },
		AllowMethods: []string{"POST"},
    AllowHeaders: []string{"Content-Type, Authorization"},
	}))
	router.POST("/apiV1/loginUser", handlers.LoginUser)
}