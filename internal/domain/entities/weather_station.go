package entities

import "gorm.io/gorm"

// WeatherStation represents a weather station with location and climate ID.
type WeatherStation struct {
	ClimateID string  `gorm:"size:255;not null"`
	Longitude float64 `gorm:"not null"`
	Latitude  float64 `gorm:"not null"`
	Name      string  `gorm:"size:255;not null"`
	gorm.Model
}
