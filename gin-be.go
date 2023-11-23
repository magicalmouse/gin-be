package main

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(200, gin.H{"id": userID, "name": "John Doe", "email": "john@example.com"})
}

func main() {
	router := gin.Default()

	UserController := &UserController{}

	router.GET("/users/:id", UserController.GetUserInfo)

	router.Run(":8080")
}
