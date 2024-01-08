package controllers

import (
	db "MyMoneyAPI/database"
	md "MyMoneyAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewClient(c *gin.Context) {
	var client md.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := md.Validate(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.DB.Create(&client)
	c.JSON(http.StatusOK, client)
}

func CardBill(c *gin.Context) {
	var card md.CreditCard
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.DB.Create(&card)
	c.JSON(http.StatusOK, card)
}
