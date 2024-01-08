package routes

import (
	ct "MyMoneyAPI/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()
	r.POST("/client", ct.NewClient)
	r.Run(":3333")
}
