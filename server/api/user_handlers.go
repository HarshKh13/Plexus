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
