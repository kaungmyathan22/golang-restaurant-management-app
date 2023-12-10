package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-restaurant-management-app/database"
	"github.com/kaungmyathan22/golang-restaurant-management-app/middlewares"
	"github.com/kaungmyathan22/golang-restaurant-management-app/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middlewares.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
