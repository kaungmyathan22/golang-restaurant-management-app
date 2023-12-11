package controllers

import "github.com/gin-gonic/gin"

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GetOrderItems"})
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GetOrderItem"})
	}
}
func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "GetOrderItemsByOrder"})
	}
}
func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "CreateOrderItem"})
	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "UpdateOrderItem"})
	}
}
