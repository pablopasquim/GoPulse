package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/GoPulse/data"
	"github.com/pablopasquim/GoPulse/models"
)

func GetItems(c *gin.Context) {

	var items []models.Item                            // slice items
	if err := data.DB.Find(&items).Error; err != nil { // find in db
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items) // return
}

func GetItemId(c *gin.Context) {

	var item models.Item
	id := c.Param("id")
	if err := data.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func CreateItem(c *gin.Context) {

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil { // read the body http request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.DB.Create(&item)       // create item
	c.JSON(http.StatusOK, item) // return
}
