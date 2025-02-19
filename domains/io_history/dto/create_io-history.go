package dto

import "mime/multipart"

type CreateIoHistoryRequest struct {
	Type    string                `form:"type" validate:"required"`
	CardId  string                `form:"card_id" validate:"required"`
	Image   *multipart.FileHeader `form:"image" validate:"required"`
	CropUrl *string               `form:"crop_url"`
}
