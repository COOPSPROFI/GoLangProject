package controllers

import (
	"github.com/COOPSPROFI/GoLangProject/pkg/configs"
	"github.com/COOPSPROFI/GoLangProject/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetAllPrinters(c *gin.Context) {
	var printer models.Printer
	configs.DB.Find(&printer)
	c.JSON(200, gin.H{
		"all printers in database": printer,
	})
}

func CreatePrinter(c *gin.Context) {
	var request struct {
		Name string
		Img  string
	}
	c.Bind(&request)

	printer := models.Printer{Name: request.Name, Img: request.Img}

	result := configs.DB.Create(&printer)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"created printer": printer,
	})
}

func GetPrinterById(c *gin.Context) {
	id := c.Param("id")
	var printer models.Printer
	configs.DB.First(&printer, id)
	c.JSON(200, gin.H{
		"get printer by id in database": printer,
	})
}

func UpdatePrinter(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Name string
		Img  string
	}
	c.Bind(&request)
	var printer models.Printer
	configs.DB.First(&printer, id)

	configs.DB.Model(&printer).Updates(models.Printer{
		Name: request.Name, Img: request.Img,
	})
	c.JSON(200, gin.H{
		"updated event": printer,
	})
}

func DeletePrinter(c *gin.Context) {
	id := c.Param("id")
	configs.DB.Delete(&models.Printer{}, id)
	c.Status(200)
}
