package main

import (
	"fmt"

	"github.com/Shoetan/db"
	"github.com/Shoetan/routers"
	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Started all over again...")
	db.InitializeDb()

	// initialize a gin engine
	r:= gin.Default()
	
	routers.RegisterUser(r)
	routers.CreateGenre(r)
	r.Run("localhost:5555")
}