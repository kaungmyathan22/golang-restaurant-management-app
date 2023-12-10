package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-restaurant-management-app/controllers"
)

func FoodRoutes(router *gin.Engine) {
	router.GET("/foods", controllers.GetFoods())
	router.GET("/foods/:food_id", controllers.GetFood())
	router.POST("/foods", controllers.CreateFood())
	router.PATCH("/foods/:food_id", controllers.UpdateFood())
}
