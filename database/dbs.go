package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
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
	database, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})
	if err != nil {
		return nil, err
	}

	// Set up connection pool
	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)

	return &Database{
		db: database,
	}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}
