package main

import (
	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/go-rest-api-starter/controllers"
	"github.com/voduybaokhanh/go-rest-api-starter/database"
)

func main() {
	database.Connect()

	r := gin.Default()

	books := r.Group("/books")
	{
		books.GET("/", controllers.GetBooks)
		books.GET("/:id", controllers.GetBook)
		books.POST("/", controllers.CreateBook)
		books.PUT("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}

	r.Run(":8080")
}
