package main

import (

	"github.com/Shoetan/db"
	"github.com/Shoetan/routers"
	"github.com/gin-gonic/gin"
)


func main() {

	db.InitializeDb()

	// initialize a gin engine
	r:= gin.Default()
	
	routers.RegisterUser(r)
	routers.CreateGenre(r)
	routers.ReturnGenre(r)
	routers.LoginUser(r)
	r.Run("localhost:5555")
}