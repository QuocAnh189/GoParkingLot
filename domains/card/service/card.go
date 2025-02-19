package service

import (
	"context"
	"goparking/domains/card/model"
	"goparking/domains/card/repository"
	"goparking/internals/libs/validation"
	"goparking/pkgs/paging"
)

type ICardService interface {
	GetCards(ctx context.Context) ([]*model.Card, *paging.Pagination, error)
	GetCardById(ctx context.Context, id string) (*model.Card, error)
	CreateCard(ctx context.Context) error
	UpdateCard(ctx context.Context) error
	DeleteCard(ctx context.Context) error
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

func (s *CardService) GetCards(ctx context.Context) ([]*model.Card, *paging.Pagination, error) {
	return nil, nil, nil
}

func (s *CardService) GetCardById(ctx context.Context, id string) (*model.Card, error) {
	return nil, nil
}

func (s *CardService) CreateCard(ctx context.Context) error {
	return nil
}

func (s *CardService) UpdateCard(ctx context.Context) error {
	return nil
}

func (s *CardService) DeleteCard(ctx context.Context) error {
	return nil
}
