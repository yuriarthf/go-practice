package controllers

import (
	"mysql-crud/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewSalesInput struct {
	Sale       []db.SaleItem `json:"sales"`
	CustomerID int64         `json:"customerId"`
}

func NewSale(c *gin.Context) {
	var ns NewSalesInput
	err := c.BindJSON(&ns)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = db.RegisterSale(ns.CustomerID, ns.Sale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// TODO: Use c.JSON for prod, since it's cheaper
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Sale Registered"})
}
