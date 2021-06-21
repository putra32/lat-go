package main

import (
	"bookstore/configs"
	"bookstore/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	configs.ConnectDataBase()

	// Router
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)

	r.Run()
}
