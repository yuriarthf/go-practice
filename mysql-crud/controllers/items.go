package controllers

import (
	"mysql-crud/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewItemInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func GetItems(c *gin.Context) {
	result, err := db.GetAllItems()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	// TODO: Use c.JSON for prod, since it's cheaper
	c.IndentedJSON(http.StatusFound, result)
}

func GetItemByName(c *gin.Context) {
	name := c.Param("name")

	result, err := db.GetItemByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	// TODO: Use c.JSON for prod, since it's cheaper
	c.IndentedJSON(http.StatusFound, result)
}

func AddItem(c *gin.Context) {
	var ni NewItemInput
	err := c.BindJSON(&ni)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	id, err := db.AddItem(ni.Name, ni.Description, ni.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// TODO: Use c.JSON for prod, since it's cheaper
	c.IndentedJSON(http.StatusCreated, gin.H{"itemId": id})
}
