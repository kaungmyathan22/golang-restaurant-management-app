package controllers

import "github.com/gin-gonic/gin"

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GetMenus"})
	}
}
func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GetMenu"})
	}
}
func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "CreateMenu"})
	}
}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "UpdateMenu"})
	}
}
