package controllers

import (
	"bookstore/configs"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	configs.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	configs.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) {
	var book models.Book

	if err := configs.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

// PATCH /books/:id
// Updata a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	configs.DB.Model(&book).Update(input)

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := configs.DB.Where("id = ? ", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	configs.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
