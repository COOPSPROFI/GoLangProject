package models

import (
	"gorm.io/gorm"
)

type Printer struct {
	gorm.Model
	Name string
	Img  string
}
