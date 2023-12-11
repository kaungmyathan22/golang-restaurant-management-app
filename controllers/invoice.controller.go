package controllers

import "github.com/gin-gonic/gin"

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GetInvoices"})
	}
}
func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "get invoice"})
	}
}
func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "CreateInvoice"})
	}
}
func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "UpdateInvoice"})
	}
}
