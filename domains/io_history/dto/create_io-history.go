package dto

import (
	"goparking/domains/io_history/model"
	"mime/multipart"
)

type CreateIoHistoryRequest struct {
	Type  string                `form:"type" validate:"required"`
	Rfid  string                `form:"rfid" validate:"required"`
	Image *multipart.FileHeader `form:"image" validate:"required"`
}

type ExitResponse struct {
	DataIn   *model.IOHistory `json:"data_in"`
	DataOut  *model.IOHistory `json:"data_out"`
	DataCard *model.Card      `json:"data_card"`
}
