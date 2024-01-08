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

func UpdateClient(c *gin.Context) {
	var client md.Client
	id := c.Params.ByName("id")
	db.DB.First(&client, id)

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

	db.DB.Model(&client).UpdateColumns(client)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Client updated!"})
}

func RemoveClient(c *gin.Context) {
	var client md.Client
	id := c.Params.ByName("id")

	db.DB.Delete(&client, id)
	c.JSON(http.StatusOK, gin.H{
		"Success": "Client deleted!"})
}
