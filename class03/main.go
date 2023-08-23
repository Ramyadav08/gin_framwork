package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "net/http"
	"fmt"
)

type Users struct {
    gorm.Model
    Name  string `json:"name"`
    Email string  `json:"email"`
}

var db *gorm.DB

var err error
var urlDns="root:@Ram81718@tcp(127.0.0.1:3306)/rekhaDB?parseTime=true"

// get all the user 
func getUsers(c *gin.Context){
	var user []Users
	db.Find(&user)
	c.JSON(http.StatusOK,user)
}

// create the new users

func createUser(c *gin.Context){
	var user Users
	// validate the json
	if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// create the user

	db.Create(&user)
    c.JSON(http.StatusCreated, user)

}

// get the user by id

func getUserByID(c *gin.Context){
	var user Users
    id := c.Param("id")
    db.First(&user, id)
    c.JSON(http.StatusOK, user)
}


// update the user 

func updateUser(c *gin.Context) {
    var user Users
    id := c.Param("id")
    if err := db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Save(&user)
    c.JSON(http.StatusOK, user)
}


// delete the user 

func deleteUser(c *gin.Context) {
    var user Users
    id := c.Param("id")
    if err := db.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    db.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


func main(){
	db,err=gorm.Open(mysql.Open(urlDns),&gorm.Config{})
	if err!=nil{
		fmt.Print(err.Error())
		panic("connection failed.......Sorry")
	}
	db.AutoMigrate(&Users{})

	r := gin.Default()

    // Define your API routes and handlers here.
    r.GET("/api/users", getUsers)
    r.GET("/api/users/:id", getUserByID)
    r.POST("/api/users", createUser)
    r.PUT("/api/users/:id", updateUser)
    r.DELETE("/api/users/:id", deleteUser)
    
    // Start the server.
    r.Run(":8080")


}