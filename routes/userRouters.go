package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/y0gesh02/go-jwt/controllers"
	"github.com/y0gesh02/go-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("users/:user_id",controller.GetUser())
	incomingRoutes.GET("/users",controller.GetUsers())
}