package repository

import (
	"context"
	dbs "goparking/database"
	"goparking/domains/io_history/model"
	"goparking/pkgs/paging"
)

type IIOHistoryRepository interface {
	GetIOHistories(ctx context.Context) ([]*model.IOHistory, *paging.Pagination, error)
	Entrance(ctx context.Context) error
	Exit(ctx context.Context) error
}

type IOHistoryRepository struct {
	db dbs.IDatabase
}

func NewIOHistoryRepository(db dbs.IDatabase) *IOHistoryRepository {
	return &IOHistoryRepository{db: db}
}

func (io *IOHistoryRepository) GetIOHistories(ctx context.Context) ([]*model.IOHistory, *paging.Pagination, error) {
	return nil, nil, nil
}

func (io *IOHistoryRepository) Entrance(ctx context.Context) error {
	return nil
}

func (io *IOHistoryRepository) Exit(ctx context.Context) error {
	return nil
}
