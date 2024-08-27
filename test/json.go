package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定義結構體 MGetBookingQ
type MGetBookingQ struct {
	Title     string    `form:"title" json:"title"`
	Sender    string    `form:"sender" json:"sender"`
	Receiver  string    `form:"receiver" json:"receiver"`
	StartTime time.Time `form:"StartTime" json:"startTime"`
	EndTime   time.Time `form:"EndTime" json:"endTime"`
}

func main() {
	r := gin.Default()

	// 路由處理表單資料
	r.POST("/submit-form", func(c *gin.Context) {
		var booking MGetBookingQ

		// 透過 ShouldBind 來綁定表單資料到 struct
		if err := c.ShouldBind(&booking); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Form data received successfully",
				"data":    booking,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 路由處理 JSON 資料
	r.POST("/submit-json", func(c *gin.Context) {
		var booking MGetBookingQ

		// 透過 ShouldBindJSON 來綁定 JSON 資料到 struct
		if err := c.ShouldBindJSON(&booking); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "JSON data received successfully",
				"data":    booking,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.Run(":8888")
}
