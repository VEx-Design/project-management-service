package gorm

import (
	gorm_model "project-management-service/external/repository/adaptors/postgres/model"

	"gorm.io/gorm"
)

func SyncDB(DB *gorm.DB) {
	DB.AutoMigrate(&gorm_model.Project{})
}
