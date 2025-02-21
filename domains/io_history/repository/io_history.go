package repository

import (
	"context"
	"goparking/configs"
	"goparking/database"
	cardModel "goparking/domains/card/model"
	"goparking/domains/io_history/dto"
	"goparking/domains/io_history/model"
	"goparking/pkgs/paging"
)

type IIOHistoryRepository interface {
	GetIOHistories(ctx context.Context, eq *dto.ListIOHistoryRequest) ([]*model.IOHistory, *paging.Pagination, error)
	ImplementEntrance(ctx context.Context, ioHistory *model.IOHistory, card *cardModel.Card) error
	ImplementExit(ctx context.Context, ioHistory *model.IOHistory, card *cardModel.Card) error
}

type IOHistoryRepository struct {
	db database.IDatabase
}

func NewIOHistoryRepository(db database.IDatabase) *IOHistoryRepository {
	return &IOHistoryRepository{db: db}
}

func (io *IOHistoryRepository) GetIOHistories(ctx context.Context, req *dto.ListIOHistoryRequest) ([]*model.IOHistory, *paging.Pagination, error) {
	ctx, cancel := context.WithTimeout(ctx, configs.DatabaseTimeout)
	defer cancel()

	query := make([]database.Query, 0)
	if req.Type != "" {
		query = append(query, database.NewQuery("type = ?", req.Type))
	}

	if req.CardType != "" {
		query = append(query, database.NewQuery("card_type = ?", req.CardType))
	}

	if req.VehicleType != "" {
		query = append(query, database.NewQuery("vehicle_type = ?", req.VehicleType))
	}

	order := "created_at DESC"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := io.db.Count(ctx, &model.IOHistory{}, &total, database.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.NewPagination(req.Page, req.Limit, total)

	var ioHistories []*model.IOHistory
	if err := io.db.Find(
		ctx,
		&ioHistories,
		database.WithQuery(query...),
		database.WithLimit(int(pagination.Size)),
		database.WithOffset(int(pagination.Skip)),
		database.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return ioHistories, pagination, nil
}

func (io *IOHistoryRepository) ImplementEntrance(ctx context.Context, ioHistory *model.IOHistory, card *cardModel.Card) error {
	handler := func() error {
		if err := io.db.Create(ctx, ioHistory); err != nil {
			return err
		}

		card.LastIOHistoryID = ioHistory.ID
		card.LastIOHistory = nil
		if err := io.db.Update(ctx, card); err != nil {
			return err
		}
		return nil
	}

	err := io.db.WithTransaction(handler)

	if err != nil {
		return err
	}
	return nil
}

func (io *IOHistoryRepository) ImplementExit(ctx context.Context, ioHistory *model.IOHistory, card *cardModel.Card) error {
	handler := func() error {
		if err := io.db.Create(ctx, ioHistory); err != nil {
			return err
		}

		card.LastIOHistoryID = ioHistory.ID
		card.LastIOHistory = nil
		if err := io.db.Update(ctx, card); err != nil {
			return err
		}
		return nil
	}

	err := io.db.WithTransaction(handler)

	if err != nil {
		return err
	}
	return nil
}
