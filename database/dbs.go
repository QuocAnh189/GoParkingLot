package database

import (
	"gorm.io/gorm"
	"time"
)

const (
	DatabaseTimeout = time.Second * 5
)

type IDatabase interface {
	GetDB() *gorm.DB
}

type Database struct {
	db *gorm.DB
}

func NewDatabase(uri string) (*Database, error) {
	return &Database{}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}
