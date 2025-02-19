package model

import (
	"github.com/google/uuid"
	"goparking/pkgs/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primary_key"`
	Email     string         `json:"email" gorm:"unique;not null;index;primary_key"`
	Name      string         `json:"name" gorm:"unique;not null;index;primary_key"`
	Password  string         `json:"password" gorm:"not null;"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))

	return nil
}

func (User) TableName() string {
	return "users"
}
