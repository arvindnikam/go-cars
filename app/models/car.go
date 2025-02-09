package models

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Make      string
	CarModel  string
	Year      int
	BodyType  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CarVariants []CarVariant
}
