package dto

import (
	"goparking/domains/card/model"
	"goparking/pkgs/paging"
)

type ListCardRequest struct {
	Search      string `json:"search,omitempty" form:"search"`
	CardType    string `json:"card_type,omitempty" form:"card_type"`
	VehicleType string `json:"vehicle_type,omitempty" form:"vehicle_type"`
	Page        int64  `json:"-" form:"page"`
	Limit       int64  `json:"-" form:"size"`
	OrderBy     string `json:"-" form:"order_by"`
	OrderDesc   bool   `json:"-" form:"order_desc"`
	TakeAll     bool   `json:"-" form:"take_all"`
}

type ListCardResponse struct {
	Cards      []*model.Card      `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}
