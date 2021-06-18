package main

import (
	"bookstore/config"
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDataBase()

	// Router
	r.GET("/books", controllers.FindBooks)

	r.Run()
}
