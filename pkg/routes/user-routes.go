package routes

import (
	"github.com/COOPSPROFI/GoLangProject/pkg/controllers"
	"github.com/COOPSPROFI/GoLangProject/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	r.GET("/logout", middlewares.RequireAuth, controllers.Logout)
}
