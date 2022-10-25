package routes

import (
	"github.com/COOPSPROFI/GoLangProject/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(r *gin.Engine) {
	r.GET("/events", controllers.GetAllEvents)
	r.POST("/events", controllers.CreateEvent)
	r.GET("/events/:id", controllers.GetEventById)
	r.PUT("/events/:id", controllers.UpdateEvent)
	r.DELETE("/events/:id", controllers.DeleteEvent)
}
