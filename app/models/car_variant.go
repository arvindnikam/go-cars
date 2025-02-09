package models

import (
	"time"

	"gorm.io/gorm"
)

type CarVariant struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	CarID        uint
	VariantCode  string
	VariantName  string
	Transmission string
	Color        string
	Engine       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`

	Car Car
}
