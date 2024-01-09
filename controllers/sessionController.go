package controllers

import (
	db "MyMoneyAPI/database"
	md "MyMoneyAPI/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	client := md.Client{}
	if err := db.DB.Where("email = ?", input.Email).First(&client).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	// Criação do token JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// Definindo as reivindicações (claims) do token
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = input.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Tempo de expiração do token

	// Assinatura do token e verificação de erros
	tokenString, err := token.SignedString([]byte("182216"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Retornando o token JWT no JSON
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
