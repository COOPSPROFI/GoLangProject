package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string
	Description string
	Img         string
	Date        time.Time
}
