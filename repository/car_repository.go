package repository

import (
	"gorm.io/gorm"
)

// CarRepository the repository of the cars
type CarRepository struct {
	db *gorm.DB
}