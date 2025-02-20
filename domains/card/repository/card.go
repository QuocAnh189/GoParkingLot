package repository

import (
	"context"
	"goparking/configs"
	"goparking/database"
	"goparking/domains/card/dto"
	"goparking/domains/card/model"
	"goparking/pkgs/paging"
)

type ICardRepository interface {
	GetCards(ctx context.Context, req *dto.ListCardRequest) ([]*model.Card, *paging.Pagination, error)
	GetCardById(ctx context.Context, id string) (*model.Card, error)
	CreatedCard(ctx context.Context, card *model.Card) error
	UpdateCard(ctx context.Context, card *model.Card) error
	DeleteCard(ctx context.Context, id string) error
}

type CardRepository struct {
	db database.IDatabase
}

func NewCardRepository(db database.IDatabase) *CardRepository {
	return &CardRepository{db: db}
}

func (c *CardRepository) GetCards(ctx context.Context, req *dto.ListCardRequest) ([]*model.Card, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Search != "" {
		query = append(query, database.NewQuery("owner_name ILIKE ?", "%"+req.Search+"%"))
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := c.db.Count(ctx, &model.Card{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	var cards []*model.Card
	if err := c.db.Find(
		ctx,
		&cards,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.Size)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return cards, pagination, nil
}

func (c *CardRepository) GetCardById(ctx context.Context, id string) (*model.Card, error) {
	var card model.Card
	if err := c.db.FindById(ctx, id, &card); err != nil {
		return nil, err
	}
	return &card, nil
}

func (c *CardRepository) CreatedCard(ctx context.Context, card *model.Card) error {
	return c.db.Create(ctx, card)
}

func (c *CardRepository) UpdateCard(ctx context.Context, card *model.Card) error {
	return c.db.Update(ctx, card)
}

func (c *CardRepository) DeleteCard(ctx context.Context, id string) error {
	card, err := c.GetCardById(ctx, id)
	if err != nil {
		return err
	}
	return c.db.Delete(ctx, card)
}
