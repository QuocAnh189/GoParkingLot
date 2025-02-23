package model

import (
	"github.com/google/uuid"
	ioHistoryModel "goparking/domains/io_history/model"
	"gorm.io/gorm"
	"time"
)

type Card struct {
	ID              string                    `json:"id" gorm:"unique;not null;index;primary_key"`
	Rfid            string                    `json:"rfid" gorm:"unique;not null;index;primary_key"`
	OwnerName       string                    `json:"owner_name" gorm:"unique;not null;index"`
	CardType        string                    `json:"card_type" gorm:"not null;index"`
	VehicleType     string                    `json:"vehicle_type" gorm:"not null;index"`
	LicensePlate    string                    `json:"license_plate" gorm:"unique;not null;index"`
	ExpiredDate     string                    `json:"expired_date" gorm:"not null;index"`
	LastIOHistoryID *string                   `json:"last_io_history_id" gorm:"default:null"` // Đổi sang con trỏ
	LastIOHistory   *ioHistoryModel.IOHistory `json:"last_io_history" gorm:"foreignKey:LastIOHistoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt       time.Time                 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time                 `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt            `json:"deleted_at" gorm:"index"`
}

func (card *Card) BeforeCreate(tx *gorm.DB) error {
	card.ID = uuid.New().String()

	return nil
}

func (Card) TableName() string {
	return "cards"
}
