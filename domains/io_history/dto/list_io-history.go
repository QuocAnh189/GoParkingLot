package dto

import (
	"goparking/domains/io_history/model"
	"goparking/pkgs/paging"
)

type ListIOHistoryRequest struct {
	Type        string `json:"type,omitempty" form:"type"`
	CardType    string `json:"card_type,omitempty" form:"card_type"`
	VehicleType string `json:"vehicle_type,omitempty" form:"vehicle_type"`
	Date        string `json:"date,omitempty" form:"date"`
	Page        int64  `json:"-" form:"page"`
	Limit       int64  `json:"-" form:"size"`
	OrderBy     string `json:"-" form:"order_by"`
	OrderDesc   bool   `json:"-" form:"order_desc"`
	TakeAll     bool   `json:"-" form:"take_all"`
}

type ListIOHistoryResponse struct {
	IOHistories []*model.IOHistory `json:"items"`
	Pagination  *paging.Pagination `json:"metadata"`
}
