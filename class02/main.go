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

func Createbook(c *gin.Context){
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	books=append(books, book)

	c.JSON(http.StatusCreated,book)

}

func GetBookById(c *gin.Context){
	id:= c.Param("id")  // parameter sent by the client, then returns that book as a response.

	for _, a:=range books{          // Loop over the list of albums, looking for
		                            // an album whose ID value matches the parameter.
		if a.ID==id{
			c.IndentedJSON(http.StatusOK, a)
            return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}



func DeleteBookById(c *gin.Context){
	id:= c.Param("id")
	 for i ,a:=range books{
		if a.ID==id{
			books = append(books[:i], books[i+1:]...)
			break
		}
	 }
	 c.JSON(http.StatusOK, gin.H{"message":"book deleted ???"})
}

func main() {
	r := gin.Default()

	r.GET("/",getbook)
	r.POST("/create",Createbook)
	r.GET("/books/:id", GetBookById)
	r.DELETE("/books/:id",DeleteBookById)
	r.Run(":5000")

}




