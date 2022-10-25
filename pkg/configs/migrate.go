package configs

import (
	"github.com/COOPSPROFI/GoLangProject/pkg/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Event{}, &models.Printer{}, &models.User{})
}
