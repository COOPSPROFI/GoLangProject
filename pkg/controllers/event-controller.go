package controllers

import (
	"time"

	"github.com/COOPSPROFI/GoLangProject/pkg/configs"
	"github.com/COOPSPROFI/GoLangProject/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	var event models.Event
	configs.DB.Find(&event)
	c.JSON(200, gin.H{
		"all events in database": event,
	})
}

func CreateEvent(c *gin.Context) {
	var request struct {
		Name        string
		Description string
		Img         string
		Date        time.Time
	}
	c.Bind(&request)

	event := models.Event{Name: request.Name, Description: request.Description, Img: request.Img, Date: request.Date}

	result := configs.DB.Create(&event)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"created event": event,
	})
}

func GetEventById(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	configs.DB.First(&event, id)
	c.JSON(200, gin.H{
		"get event by id in database": event,
	})
}

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Name        string
		Description string
		Img         string
		Date        time.Time
	}
	c.Bind(&request)
	var event models.Event
	configs.DB.First(&event, id)

	configs.DB.Model(&event).Updates(models.Event{
		Name: request.Name, Description: request.Description, Img: request.Img, Date: request.Date,
	})
	c.JSON(200, gin.H{
		"updated event": event,
	})
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	configs.DB.Delete(&models.Event{}, id)
	c.Status(200)
}
