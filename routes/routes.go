package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/GoPulse/controllers"
)

func HandleRequest(r *gin.Engine) { // endpoints
	r.GET("/api/items", controllers.GetItems)
	r.GET("/api/items/:id", controllers.GetItemId)
	r.POST("/api/items", controllers.CreateItem)
}
