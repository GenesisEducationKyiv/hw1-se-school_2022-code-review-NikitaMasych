package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PresentRateJSON(c *gin.Context, rate float64, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"description": "Invalid status value"})
	} else {
		c.JSON(http.StatusOK, gin.H{"description": rate})
	}
}

func PresentUserConflictJSON(c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{"description": "User is already subscribed"})
}

func PresentUserSubscriptionJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"description": "User successfully subscribed"})
}

func PresentEmailsSentJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"description": "Emails has been sent"})
}
