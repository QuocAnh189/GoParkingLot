package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type IOHistory struct {
	ID          string         `json:"id" gorm:"unique;not null;index;primary_key"`
	Type        string         `json:"type" gorm:"not null"`
	CardID      string         `json:"card_id" gorm:"not null"`
	ImageUrl    string         `json:"image_url" gorm:"not null"`
	CardType    string         `json:"card_type" gorm:"not null"`
	VehicleType string         `json:"vehicle_type" gorm:"not null"`
	CropUrl     string         `json:"crop_url" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type Card struct {
	ID           string `json:"id"`
	OwnerName    string `json:"owner_name" `
	CardType     string `json:"card_type"`
	VehicleType  string `json:"vehicle_type" `
	LicensePlate string `json:"license_plate"`
	ExpiredDate  string `json:"expired_date"`
}

func (io *IOHistory) BeforeCreate(tx *gorm.DB) error {
	io.ID = uuid.New().String()

	return nil
}

func (IOHistory) TableName() string {
	return "io_histories"
}
