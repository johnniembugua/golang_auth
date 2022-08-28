package routes

import (
	middleware "github.com/johnniembugua/golang-jwt-project/midlleware"

	controller "github.com/johnniembugua/golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
