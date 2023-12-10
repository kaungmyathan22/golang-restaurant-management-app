package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-restaurant-management-app/controllers"
)

func OrderItemsRoutes(router *gin.Engine) {
	router.GET("/orderItems", controllers.GetOrderItems())
	router.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	router.GET("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
	router.POST("/orderItems", controllers.CreateOrderItem())
	router.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}
