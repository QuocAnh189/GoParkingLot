package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Card struct {
	ID           string         `json:"id" gorm:"unique;not null;index;primary_key"`
	Rfid         string         `json:"rfid" gorm:"unique;not null;index;primary_key"`
	OwnerName    string         `json:"owner_name" gorm:"unique;not null;index"`
	CardType     string         `json:"card_type" gorm:"not null;index"`
	VehicleType  string         `json:"vehicle_type" gorm:"not null;index"`
	LicensePlate string         `json:"license_plate" gorm:"unique;not null;index"`
	ExpiredDate  string         `json:"expired_date" gorm:"not null;index"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (card *Card) BeforeCreate(tx *gorm.DB) error {
	card.ID = uuid.New().String()

	return nil
}

func (Card) TableName() string {
	return "cards"
}
