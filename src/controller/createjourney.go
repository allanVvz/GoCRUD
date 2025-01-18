package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateJourney : CreateJourney
func CreateJourney(c *gin.context) {
	fmt.Println("CreateJourney")
	c.JSON(200, gin.H{
		"message": "CreateJourney",
	})
}
