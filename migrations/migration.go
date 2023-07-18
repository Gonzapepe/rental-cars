package migrations

import (
	"github.com/Gonzapepe/cars-rental/helper"
	"github.com/Gonzapepe/cars-rental/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Car{})
	helper.ErrorPanic(err)

	err = db.AutoMigrate(&model.Make{})
	helper.ErrorPanic(err)

	return nil
}