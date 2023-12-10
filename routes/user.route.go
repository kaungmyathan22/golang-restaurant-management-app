package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-restaurant-management-app/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers())
	router.GET("/users/:userId", controllers.GetUser())
	router.POST("/users/signup", controllers.SignUp())
	router.POST("/users/login", controllers.Login())
}
