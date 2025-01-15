package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ListJourneys(c *gin.Context) {
	fmt.Println("FindJourneys")
	c.JSON(200, gin.H{
		"message": "FindJourneys",
	})
}

func List2wJourneys(c *gin.Context) {
	fmt.Println("FindJourney")
	c.JSON(200, gin.H{
		"message": "Find 2 weeks Journey",
	})
}