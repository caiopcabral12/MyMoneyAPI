package controllers

import (
	db "MyMoneyAPI/database"
	md "MyMoneyAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBill(c *gin.Context) {
	var bill md.Payable

	id := c.Params.ByName("id")
	client := md.Client{}

	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := db.DB.First(&client, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User Not Found!"})
		return
	}

	bill.ClientID = client.ID

	db.DB.Create(&bill)

	db.DB.Preload("Client").First(&bill, bill.ID)

	c.JSON(http.StatusOK, bill)
}

func UpdateBill(c *gin.Context) {
	var bill md.Payable
	id := c.Params.ByName("id")
	db.DB.First(&bill, id)

	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.DB.Model(&bill).UpdateColumns(bill)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Bill updated!"})
}

func RemoveBill(c *gin.Context) {
	var bill md.Payable
	id := c.Params.ByName("id")

	db.DB.Delete(&bill, id)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Bill deleted!"})
}
