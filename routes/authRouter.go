package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manavgupta457/go-jwt-project/controllers"
)

func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}