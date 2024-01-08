package main

import (
	ct "MyMoneyAPI/controllers"
	db "MyMoneyAPI/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db.DbConnect()
	r := gin.Default()
	r.POST("/client", ct.NewClient)
	r.POST("/client/id", ct.CardBill)

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3333")

}
