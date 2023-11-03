package routers

import (
	"github.com/Shoetan/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUser (router *gin.Engine) {
	router.POST("/registerUser", handlers.RegisterUser)
}