package controllers

import (
	db "MyMoneyAPI/database"
	md "MyMoneyAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSaving(c *gin.Context) {
	var saving md.Savings

	id := c.Params.ByName("id")
	client := md.Client{}

	if err := c.ShouldBindJSON(&saving); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User Not Found!"})
		return
	}

	saving.ClientID = client.ID

	db.DB.Create(&saving)

	db.DB.Preload("Client").First(&saving, saving.ID)

	c.JSON(http.StatusOK, saving)
}

func UpdateSaving(c *gin.Context) {
	var saving md.Savings
	id := c.Params.ByName("id")
	db.DB.First(&saving, id)

	if err := c.ShouldBindJSON(&saving); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.DB.Model(&saving).UpdateColumns(saving)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Saving updated!"})
}

func RemoveSaving(c *gin.Context) {
	var saving md.Savings
	id := c.Params.ByName("id")

	db.DB.Delete(&saving, id)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Saving deleted!"})
}
