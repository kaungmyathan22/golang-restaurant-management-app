package controllers

import "github.com/gin-gonic/gin"

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GetTables"})
	}
}
func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GetTable"})
	}
}
func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "CreateTable"})
	}
}
func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "UpdateTable"})
	}
}
