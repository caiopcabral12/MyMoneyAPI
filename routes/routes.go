package routes

import (
	ct "MyMoneyAPI/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

	r.POST("/login", ct.Login)

	r.POST("/client", ct.NewClient)
	r.PATCH("/client/:id", ct.UpdateClient)
	r.DELETE("/client/:id", ct.RemoveClient)

	r.POST("/card/:id", ct.CardBill)
	r.PATCH("/card/:id", ct.UpdateCard)
	r.DELETE("/card/:id", ct.RemoveCard)

	r.POST("/bills/:id", ct.CreateBill)
	r.PATCH("/bills/:id", ct.UpdateBill)
	r.DELETE("/bills/:id", ct.RemoveBill)

	r.POST("/receivable/:id", ct.CreateReceivable)
	r.PATCH("/receivable/:id", ct.UpdateReceivable)
	r.DELETE("/receivable/:id", ct.RemoveReceivable)

	r.POST("/savings/:id", ct.CreateSaving)
	r.PATCH("/savings/:id", ct.UpdateSaving)
	r.DELETE("/savings/:id", ct.RemoveSaving)

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

	r.Run(":3333")

}
