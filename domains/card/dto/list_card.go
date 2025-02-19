package dto

import (
	"goparking/domains/card/model"
	"goparking/pkgs/paging"
)

type ListCardRequest struct {
	Search    string `json:"name,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}

type ListCardResponse struct {
	Coupon     []*model.Card      `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}
