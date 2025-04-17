package api

import (
	"log"
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var err error
	database, err := database.DatabaseConnection()
	if err != nil {
		log.Printf("Failed to connect to Database")
	}

	var users []models.Users
	result := database.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func CreateUser(c *gin.Context) {
	var err error
	database, err := database.DatabaseConnection()
	if err != nil {
		log.Printf("Failed to connect to Database")
	}

	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Input Data"})
		return
	}

	result := database.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created Successfully"})

}

func GetUserByUsernameAndPassword(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var err error
	database, err := database.DatabaseConnection()
	if err != nil {
		log.Printf("Failed to connect to Database")
	}

	var user models.Users
	result := database.Where("username=?, password=?", username, password).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
