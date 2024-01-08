package controllers

import (
	db "MyMoneyAPI/database"
	md "MyMoneyAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CardBill(c *gin.Context) {
	var card md.CreditCard

	id := c.Params.ByName("id")
	client := md.Client{}

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User Not Found!"})
		return
	}

	card.ClientID = client.ID

	db.DB.Create(&card)

	db.DB.Preload("Client").First(&card, card.ID)

	c.JSON(http.StatusOK, card)
}

func UpdateCard(c *gin.Context) {
	var card md.CreditCard
	id := c.Params.ByName("id")
	db.DB.First(&card, id)

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.DB.Model(&card).UpdateColumns(card)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Card updated!"})
}

func RemoveCard(c *gin.Context) {
	var card md.CreditCard
	id := c.Params.ByName("id")

	db.DB.Delete(&card, id)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Card deleted!"})
}
