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
	"goparking/proto/gen/pb_detects"
	"io"
	"mime/multipart"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	cardRepo cardRepo.ICardRepository,
) *IOHistoryService {
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

	layout := "2006-01-02"
	expiredTime, err := time.Parse(layout, card.ExpiredDate)
	if err != nil {
		return nil, errors.New("invalid expired date format")
	}

	if expiredTime.Before(time.Now()) {
		return nil, errors.New("expired date")
	}

	if card.LastIOHistory != nil && card.LastIOHistory.Type == req.Type {
		return nil, errors.New("same type IN")
	}

	var ioHistory *model.IOHistory
	utils.MapStruct(&ioHistory, req)

	listPlates, cropImgUrl, err := DetectPlate(req.Image)
	if err != nil {
		return nil, err
	}

	if len(listPlates) == 0 {
		return nil, errors.New("no plate")
	}

	if listPlates[0] != card.LicensePlate {
		return nil, errors.New("license_plate not right")
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

	ioHistory.ImageUrl = imageUrlUpload
	ioHistory.CropUrl = cropImgUrl
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

	layout := "2006-01-02"
	expiredTime, err := time.Parse(layout, card.ExpiredDate)
	if err != nil {
		return nil, nil, nil, errors.New("invalid expired date format")
	}

	if expiredTime.Before(time.Now()) {
		return nil, nil, nil, errors.New("expired date")
	}

	if card.LastIOHistory == nil {
		return nil, nil, nil, errors.New("the car is not in the yard")
	}

	if card.LastIOHistory == nil {
		return nil, nil, nil, errors.New("license_plate not right")
	}

	if card.LastIOHistory != nil && card.LastIOHistory.Type == req.Type {
		return nil, nil, nil, errors.New("same type OUT")
	}

	var inHistory = card.LastIOHistory
	var outHistory *model.IOHistory
	utils.MapStruct(&outHistory, req)

	listPlates, cropImgUrl, err := DetectPlate(req.Image)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(listPlates) == 0 {
		return nil, nil, nil, errors.New("no plate")
	}

	if listPlates[0] != card.LicensePlate {
		return nil, nil, nil, errors.New("license_plate not right")
	}

	var imageUrlUpload = ""
	if req.Image != nil {
		imageURL, err := io.minioClient.UploadFile(ctx, req.Image, "io_histories")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return nil, nil, nil, err
		}
		imageUrlUpload = imageURL
	}

	outHistory.ImageUrl = imageUrlUpload
	outHistory.CropUrl = cropImgUrl
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

func DetectPlate(image *multipart.FileHeader) ([]string, string, error) {
	file, err := image.Open()
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	imageData, err := io.ReadAll(file)
	if err != nil {
		return nil, "", err
	}

	// conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", conf.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("parking.plate_detector:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, "", err
	}
	defer conn.Close()

	//// Load TLS Cert - Deploy Https
	//creds := credentials.NewClientTLSFromCert(nil, "")
	//
	//// Kết nối tới gRPC Server qua HTTPS (nginx reverse proxy)
	//conn, err := grpc.NewClient("license-detect.duckdns.org:9443", grpc.WithTransportCredentials(creds))
	//if err != nil {
	//	return nil, "", err
	//}
	//defer conn.Close()

	client := pb_detects.NewPlateDetectionClient(conn)

	req := &pb_detects.PlateRequest{Image: imageData}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	response, err := client.Detect(ctxTimeout, req)
	if err != nil {
		return nil, "", err
	}

	return response.LicensePlateDetect, response.CropImgUrl, nil
}
