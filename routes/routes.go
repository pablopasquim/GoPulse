package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/GoPulse/controllers"
)

func HandleRequest(r *gin.Engine) { // endpoints
	r.GET("/api/items", controllers.GetItems)
	r.POST("/api/create", controllers.CreateItem)
}
