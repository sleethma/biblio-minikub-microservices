package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to my library"})
	})

	engine.GET("/api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, AllBooks())
	})

	engine.GET("/api/books/:isbn", func(ctx *gin.Context) {
		isbn := ctx.Params.ByName("isbn")
		book, found := GetBook(isbn)
		if found {
			ctx.JSON(http.StatusOK, book)
		} else {
			ctx.AbortWithStatus(http.StatusNotFound)
		}
	})

	engine.POST("/api/books", func(ctx *gin.Context) {
		var book Book
		if ctx.BindJSON(&book) == nil {
			isbn, created := CreateBook(book)
			if created {
				ctx.Header("Location", "/books/api/"+isbn)
				ctx.Status(http.StatusCreated)
			} else {
				ctx.Status(http.StatusConflict)
			}
		}
	})

	engine.PUT("api/books/:isbn", func(ctx *gin.Context) {
		var book Book
		isbn := ctx.Params.ByName("isbn")
		if ctx.BindJSON(&book) == nil {
			exists := UpdateBook(isbn, book)
			if exists {
				ctx.Status(http.StatusOK)
			} else {
				ctx.Status(http.StatusConflict)
			}
		}
	})

	engine.DELETE("/api/books/:isbn", func(ctx *gin.Context) {
		DeleteBook(ctx.Params.ByName("isbn"))
		ctx.Status(http.StatusOK)
	})

	engine.Run(port())

}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
