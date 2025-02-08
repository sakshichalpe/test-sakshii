package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// EmailRequest represents an email send request
type EmailRequest struct {
	To      string `json:"to" binding:"required,email"`
	From    string `json:"from" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

// EmailStatistics stores email API usage statistics
type EmailStatistics struct {
	TotalEmailsSent int
	TotalFailures   int
	Mutex           sync.Mutex
}

var stats = EmailStatistics{}

func sendEmail(c *gin.Context) {
	var emailReq EmailRequest
	if err := c.ShouldBindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulate email warming up limit
	stats.Mutex.Lock()
	if stats.TotalEmailsSent < 5 {
		stats.TotalEmailsSent++
		stats.Mutex.Unlock()
		c.JSON(http.StatusOK, gin.H{"message": "Email queued successfully"})
	} else {
		stats.TotalFailures++
		stats.Mutex.Unlock()
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Email sending limit reached due to warming up phase"})
	}
}

func getStatistics(c *gin.Context) {
	stats.Mutex.Lock()
	defer stats.Mutex.Unlock()
	c.JSON(http.StatusOK, stats)
}

func main() {
	r := gin.Default()
	r.POST("/v1/email/send", sendEmail)
	r.GET("/v1/email/statistics", getStatistics)

	r.Run(":8080") // Run server on port 8080
}
