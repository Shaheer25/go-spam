package main

import (
	"strconv"

	_ "github.com/Shaheer25/rest-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Book represents a book entity.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

// @title My Book API
// @version 1.0
// @description API for managing books
// @BasePath /api/v1
func main() {
	r := gin.Default()

	// Swagger handler
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1/books")
	{
		v1.GET("/", GetBooks)
		v1.GET("/:id", GetBook)
		v1.POST("/", CreateBook)
		v1.PUT("/:id", UpdateBook)
		v1.DELETE("/:id", DeleteBook)
	}

	r.Run(":8080")
}

// GetBooks returns the list of all books.
// @Summary Get all books
// @Description Get all books
// @Produce json
// @Success 200 {array} Book
// @Router /api/v1/books/ [get]
func GetBooks(c *gin.Context) {
	c.JSON(200, books)
}

// GetBook returns a specific book by ID.
// @Summary Get a book by ID
// @Description Get a specific book by ID
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} Book
// @Router /api/v1/books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == id {
			c.JSON(200, book)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Book not found"})
}

// CreateBook creates a new book.
// @Summary Create a book
// @Description Create a new book
// @Accept json
// @Produce json
// @Param book body Book true "Book object"
// @Success 201 {object} Book
// @Router /api/v1/books/ [post]
func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book.ID = "b" + strconv.Itoa(len(books)+1)
	books = append(books, book)

	c.JSON(201, book)
}

// UpdateBook updates a book by ID.
// @Summary Update a book
// @Description Update a specific book by ID
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body Book true "Book object"
// @Success 200 {object} Book
// @Router /api/v1/books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var UpdatedBook Book
	if err := c.ShouldBindJSON(&UpdatedBook); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i].Title = UpdatedBook.Title
			books[i].Author = UpdatedBook.Author
			c.JSON(200, books[i])
			return
		}
	}

	c.JSON(404, gin.H{"error": "Book not found"})
}

// DeleteBook deletes a book by ID.
// @Summary Delete a book
// @Description Delete a specific book by ID
// @Produce json
// @Param id path string true "Book ID"
// @Success 204
// @Router /api/v1/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.Status(204)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Book not found"})
}
