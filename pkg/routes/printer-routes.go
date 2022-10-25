package routes

import (
	"github.com/COOPSPROFI/GoLangProject/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterPrinterRoutes(r *gin.Engine) {
	r.GET("/printers", controllers.GetAllPrinters)
	r.POST("/printers", controllers.CreatePrinter)
	r.GET("/printers/:id", controllers.GetPrinterById)
	r.PUT("/printers/:id", controllers.UpdatePrinter)
	r.DELETE("/printers/:id", controllers.DeletePrinter)
}
