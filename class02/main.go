package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func getbook(c *gin.Context){
	c.JSON(http.StatusOK,books)
}

func main() {
	r := gin.Default()

	r.GET("/",getbook)
	r.Run(":5000")

}




