package controllers

import (
	"bookstore/config"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
