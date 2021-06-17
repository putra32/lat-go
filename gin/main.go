package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Router biasa
	router.GET("/user/", User)

	// Router group
	v1 := router.Group("/v1")
	{
		v1.GET("/user", User)
	}

	v2 := router.Group("v2")
	{
		v2.GET("/user", NewUser)
	}

	// Router static
	router.Static("/assets", "./assets")

	// Router dengan template html
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"judul": "Response dengan Output HTML",
		})
	})

	// Router dengan string/text
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Cara menggunakan output string Gin")
	})

	// Router dengan type json
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ini Pesan JSON",
			"status":  http.StatusOK,
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		// Dengan Path
		name := c.Param("name")
		// Dengan Query
		address := c.Query("address")

		c.String(http.StatusOK, "Hello %s Alamat %s ", name, address)
	})

	router.Run()
}

func User(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Dari V1"))
}

func NewUser(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Dari V2"))
}
