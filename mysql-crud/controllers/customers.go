package controllers

import (
	"mysql-crud/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewCustomerInput struct {
	Name    string `json:"name"`
	Age     uint8  `json:"age"`
	Address string `json:"address"`
}

func GetCustomers(c *gin.Context) {
	result, err := db.GetAllCustomers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	// TODO: Use c.JSON for prod, since it's cheaper
	c.IndentedJSON(http.StatusFound, result)
}

func GetCustomerById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result, err := db.GetCustomerById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	// TODO: Use c.JSON for prod, since it's cheaper
	c.IndentedJSON(http.StatusFound, result)
}

func NewCustomer(c *gin.Context) {
	var nc NewCustomerInput
	err := c.BindJSON(&nc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	id, err := db.NewCustomer(nc.Name, nc.Age, nc.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// TODO: Use c.JSON for prod, since it's cheaper
	c.IndentedJSON(http.StatusCreated, gin.H{"customerId": id})
}
