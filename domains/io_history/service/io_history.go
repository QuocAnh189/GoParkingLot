package service

import (
	"context"
	"goparking/domains/io_history/model"
	"goparking/domains/io_history/repository"
	"goparking/internals/libs/validation"
	"goparking/pkgs/paging"
)

type IIOHistoryService interface {
	GetIOHistories(ctx context.Context) ([]*model.IOHistory, *paging.Pagination, error)
	Entrance(ctx context.Context) error
	Exit(ctx context.Context) error
}

type IOHistoryService struct {
	validator     validation.Validation
	ioHistoryRepo repository.IIOHistoryRepository
}

func NewIOHistoryService(validator validation.Validation, ioHistoryRepo repository.IIOHistoryRepository) *IOHistoryService {
	return &IOHistoryService{
		validator:     validator,
		ioHistoryRepo: ioHistoryRepo,
	}
}

func (io *IOHistoryService) GetIOHistories(ctx context.Context) ([]*model.IOHistory, *paging.Pagination, error) {
	return nil, nil, nil
}

func (io *IOHistoryService) Entrance(ctx context.Context) error {
	return nil
}

func (io *IOHistoryService) Exit(ctx context.Context) error {
	return nil
}
