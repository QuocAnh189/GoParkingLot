package dto

import "mime/multipart"

type SignUpRequest struct {
	Email    string                `form:"email" validate:"required,email"`
	Name     string                `form:"name" validate:"required"`
	Avatar   *multipart.FileHeader `form:"avatar"`
	Password string                `form:"password" validate:"required"`
}

type SignUpResponse struct {
	AccessToken  string `json:"accessToken" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}
