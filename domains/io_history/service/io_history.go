package service

import (
	"context"
	"errors"
	cardRepo "goparking/domains/card/repository"
	"goparking/domains/io_history/dto"
	"goparking/domains/io_history/model"
	"goparking/domains/io_history/repository"
	"goparking/internals/libs/logger"
	"goparking/internals/libs/validation"
	"goparking/pkgs/minio"
	"goparking/pkgs/paging"
	"goparking/pkgs/utils"
)

type IIOHistoryService interface {
	GetIOHistories(ctx context.Context, req *dto.ListIOHistoryRequest) ([]*model.IOHistory, *paging.Pagination, error)
	Entrance(ctx context.Context, req *dto.CreateIoHistoryRequest) (*model.IOHistory, error)
	Exit(ctx context.Context, req *dto.CreateIoHistoryRequest) (*model.IOHistory, *model.IOHistory, *model.Card, error)
}

type IOHistoryService struct {
	validator     validation.Validation
	ioHistoryRepo repository.IIOHistoryRepository
	cardRepo      cardRepo.ICardRepository
	minioClient   *minio.MinioClient
}

func NewIOHistoryService(
	validator validation.Validation,
	ioHistoryRepo repository.IIOHistoryRepository,
	minioClient *minio.MinioClient,
	cardRepo cardRepo.ICardRepository) *IOHistoryService {
	return &IOHistoryService{
		validator:     validator,
		ioHistoryRepo: ioHistoryRepo,
		cardRepo:      cardRepo,
		minioClient:   minioClient,
	}
}

func (io *IOHistoryService) GetIOHistories(ctx context.Context, req *dto.ListIOHistoryRequest) ([]*model.IOHistory, *paging.Pagination, error) {
	ioHistories, pagination, err := io.ioHistoryRepo.GetIOHistories(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return ioHistories, pagination, nil
}

func (io *IOHistoryService) Entrance(ctx context.Context, req *dto.CreateIoHistoryRequest) (*model.IOHistory, error) {
	if err := io.validator.ValidateStruct(req); err != nil {
		return nil, err
	}

	card, err := io.cardRepo.GetCardByRFID(ctx, req.Rfid)
	if err != nil {
		return nil, errors.New("invalid card")
	}

	if card.LastIOHistory != nil && card.LastIOHistory.Type == req.Type {
		return nil, errors.New("same type IN")
	}

	var imageUrlUpload = ""
	if req.Image != nil {
		imageURL, err := io.minioClient.UploadFile(ctx, req.Image, "io_histories")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return nil, err
		}
		imageUrlUpload = imageURL
	}

	var ioHistory *model.IOHistory
	utils.MapStruct(&ioHistory, req)
	ioHistory.ImageUrl = imageUrlUpload
	ioHistory.CropUrl = "http:hahaha"
	ioHistory.CardID = card.ID
	ioHistory.CardType = card.CardType
	ioHistory.VehicleType = card.VehicleType

	if err := io.ioHistoryRepo.ImplementExit(ctx, ioHistory, card); err != nil {
		return nil, err
	}
	return ioHistory, nil
}

func (io *IOHistoryService) Exit(ctx context.Context, req *dto.CreateIoHistoryRequest) (*model.IOHistory, *model.IOHistory, *model.Card, error) {
	if err := io.validator.ValidateStruct(req); err != nil {
		return nil, nil, nil, err
	}

	card, err := io.cardRepo.GetCardByRFID(ctx, req.Rfid)
	if err != nil {
		return nil, nil, nil, errors.New("invalid card")
	}

	if card.LastIOHistory == nil {
		return nil, nil, nil, errors.New("the car is not in the yard")
	}

	if card.LastIOHistory != nil && card.LastIOHistory.Type == req.Type {
		return nil, nil, nil, errors.New("same type OUT")
	}

	var inHistory = card.LastIOHistory

	var imageUrlUpload = ""
	if req.Image != nil {
		imageURL, err := io.minioClient.UploadFile(ctx, req.Image, "io_histories")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return nil, nil, nil, err
		}
		imageUrlUpload = imageURL
	}

	var outHistory *model.IOHistory
	utils.MapStruct(&outHistory, req)
	outHistory.ImageUrl = imageUrlUpload
	outHistory.CropUrl = "http:hahaha"
	outHistory.CardID = card.ID
	outHistory.CardType = card.CardType
	outHistory.VehicleType = card.VehicleType

	if err := io.ioHistoryRepo.ImplementEntrance(ctx, outHistory, card); err != nil {
		return nil, nil, nil, err
	}

	var dataCard model.Card
	utils.MapStruct(&dataCard, card)
	return outHistory, inHistory, &dataCard, nil
}
