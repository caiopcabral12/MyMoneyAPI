package controllers

import (
	db "MyMoneyAPI/database"
	md "MyMoneyAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateReceivable(c *gin.Context) {
	var receives md.Receivable

	id := c.Params.ByName("id")
	client := md.Client{}

	if err := c.ShouldBindJSON(&receives); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User Not Found!"})
		return
	}

	receives.ClientID = client.ID

	db.DB.Create(&receives)

	db.DB.Preload("Client").First(&receives, receives.ID)

	c.JSON(http.StatusOK, receives)
}

func UpdateReceivable(c *gin.Context) {
	var receives md.Receivable
	id := c.Params.ByName("id")
	db.DB.First(&receives, id)

	if err := c.ShouldBindJSON(&receives); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.DB.Model(&receives).UpdateColumns(receives)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Receive updated!"})
}

func RemoveReceivable(c *gin.Context) {
	var receives md.Receivable
	id := c.Params.ByName("id")

	db.DB.Delete(&receives, id)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Receive deleted!"})
}
