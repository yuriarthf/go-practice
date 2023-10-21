package main

import (
	"log"
	"mysql-crud/controllers"
	"mysql-crud/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func loadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading config file")
	}
}

func main() {
	// Load env variables from .env
	loadEnv()

	// Instantiate SQL database
	db.ConfigMySQL()
	defer db.Close()

	// Instantiate and configure Gin server
	r := gin.Default()

	// Define endpoints
	/// GET Methods
	r.GET("/customers", controllers.GetCustomers)
	r.GET("/customers/id/:id", controllers.GetCustomerById)
	r.GET("/items", controllers.GetItems)
	r.GET("/items/name/:name", controllers.GetItemByName)
	//r.GET("/sales", getSales)

	/// POST METHODS
	r.POST("/newCustomer", controllers.NewCustomer)

	log.Fatal(r.Run())
}
