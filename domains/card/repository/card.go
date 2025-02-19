package repository

import (
	"context"
	dbs "goparking/database"
	"goparking/domains/card/model"
	"goparking/pkgs/paging"
)

type ICardRepository interface {
	GetCards(ctx context.Context) ([]*model.Card, *paging.Pagination, error)
	GetCardById(ctx context.Context, id string) (*model.Card, error)
	CreatedCard(ctx context.Context) error
	UpdateCard(ctx context.Context) error
	DeleteCard(ctx context.Context) error
}

type CardRepository struct {
	db dbs.IDatabase
}

func NewCardRepository(db dbs.IDatabase) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) GetCards(ctx context.Context) ([]*model.Card, *paging.Pagination, error) {
	return nil, nil, nil
}

func (r *CardRepository) GetCardById(ctx context.Context, id string) (*model.Card, error) {
	return nil, nil
}

func (r *CardRepository) CreatedCard(ctx context.Context) error {
	return nil
}

func (r *CardRepository) UpdateCard(ctx context.Context) error {
	return nil
}

func (r *CardRepository) DeleteCard(ctx context.Context) error {
	return nil
}
