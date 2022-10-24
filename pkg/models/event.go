package models

import (
	"time"

	"github.com/COOPSPROFI/GoLangProject/pkg/config"
	"github.com/jinzhu/gorm"
)

var (
	database *gorm.DB
)

type Event struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Img         string    `json:"img"`
	Date        time.Time `json:"date"`
}

func Init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Event{})
}

func GetEvents() []Event {
	var Events []Event
	database.Find(&Events)
	return Events
}

func (e *Event) CreateEvent() *Event {
	database.NewRecord(e)
	database.Create(&e)
	return e
}

func GetEventById(Id int64) (*Event, *gorm.DB) {
	var getEvent Event
	db := database.Where("Id=?", Id).Find(&getEvent)
	return &getEvent, db
}

func DeleteEvent(ID int64) Event {
	var event Event
	database.Where("ID=?", ID).Delete(event)
	return event
}
