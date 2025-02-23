package service

import (
	"context"
	"goparking/domains/card/dto"
	"goparking/domains/card/model"
	"goparking/domains/card/repository"
	"goparking/internals/libs/logger"
	"goparking/internals/libs/validation"
	"goparking/pkgs/paging"
	"goparking/pkgs/utils"
)

type ICardService interface {
	GetCards(ctx context.Context, req *dto.ListCardRequest) ([]*model.Card, *paging.Pagination, error)
	GetCardById(ctx context.Context, id string) (*model.Card, error)
	CreateCard(ctx context.Context, req *dto.CreateCardRequest) error
	UpdateCard(ctx context.Context, id string, req *dto.UpdateCardRequest) error
	DeleteCard(ctx context.Context, id string) error
}

type CardService struct {
	validator validation.Validation
	cardRepo  repository.ICardRepository
}

func NewCardService(validator validation.Validation, cardRepo repository.ICardRepository) *CardService {
	return &CardService{
		validator: validator,
		cardRepo:  cardRepo,
	}
}

func (c *CardService) GetCards(ctx context.Context, req *dto.ListCardRequest) ([]*model.Card, *paging.Pagination, error) {
	cards, pagination, err := c.cardRepo.GetCards(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	return cards, pagination, nil
}

func (c *CardService) GetCardById(ctx context.Context, id string) (*model.Card, error) {
	card, err := c.cardRepo.GetCardById(ctx, id)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (c *CardService) CreateCard(ctx context.Context, req *dto.CreateCardRequest) error {
	if err := c.validator.ValidateStruct(req); err != nil {
		return err
	}

	var card model.Card
	utils.MapStruct(&card, req)

	err := c.cardRepo.CreatedCard(ctx, &card)
	if err != nil {
		logger.Errorf("Create fail, error: %s", err)
		return err
	}
	return nil
}

func (c *CardService) UpdateCard(ctx context.Context, id string, req *dto.UpdateCardRequest) error {
	if err := c.validator.ValidateStruct(req); err != nil {
		return err
	}

	card, err := c.cardRepo.GetCardById(ctx, id)
	if err != nil {
		logger.Errorf("Get fail, error: %s", err)
		return err
	}

	utils.MapStruct(card, req)
	//if req.LastIOHistoryID == nil {
	//	card.LastIOHistoryID = nil
	//}
	err = c.cardRepo.UpdateCard(ctx, card)
	if err != nil {
		logger.Errorf("Update fail, id: %s, error: %s", id, err)
		return err
	}

	return nil
}

func (c *CardService) DeleteCard(ctx context.Context, id string) error {
	err := c.cardRepo.DeleteCard(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
