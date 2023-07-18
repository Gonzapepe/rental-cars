package model

import "time"

// Car struct of the Car table
type Car struct {
	Id int `json:"id,omitempty" gorm:"primaryKey"`
	Model string `json:"model,omitempty" gorm:"column:model"`
	Drive string `json:"drive,omitempty" gorm:"column:drive"`
	FuelType string `json:"fuel_type,omitempty" gorm:"column:fuel_type"`
	Year int `json:"year,omitempty" gorm:"column:year"`
	Transmission string `json:"transmission,omitempty" gorm:"column:transmission"`
	MakeID int `json:"make_id,omitempty" gorm:"column:make_id"`
	Make Make `json:"make,omitempty" gorm:"foreignkey:makeID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName specifies the database table name for the Car Model.
func (Car) TableName() string {
	return "cars"
}

// Make struct of the Maker of the car
type Make struct {
	Id int `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name,omitempty" gorm:"column:name"`
	Cars []Car `gorm:"foreignkey:MakeID"`
}

// TableName specifies the database table name for the Make Model.
func (Make) TableName() string {
	return "makers"
}