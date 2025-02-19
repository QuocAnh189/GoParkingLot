package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type IOHistory struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primary_key"`
	Type      string         `json:"type" gorm:"not null"`
	CardId    string         `json:"card_id" gorm:"not null"`
	ImageUrl  string         `json:"image_url" gorm:"not null"`
	CropUrl   string         `json:"crop_url" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (io *IOHistory) BeforeCreate(tx *gorm.DB) error {
	io.ID = uuid.New().String()

	return nil
}

func (IOHistory) TableName() string {
	return "io_histories"
}
