package controllers

import "github.com/gin-gonic/gin"

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GetFoods"})
	}
}
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "get food"})
	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "CreateFood"})
	}
}
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "UpdateFood"})
	}
}
