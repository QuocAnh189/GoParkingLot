package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"goparking/domains/auth/dto"
	"goparking/domains/auth/model"
	"goparking/domains/auth/repository"
	"goparking/internals/libs/logger"
	"goparking/internals/libs/validation"
	"goparking/pkgs/jwt"
	"goparking/pkgs/minio"
	"goparking/pkgs/utils"
)

type IUserService interface {
	SignIn(ctx context.Context, req *dto.SignInRequest) (string, string, error)
	SignUp(ctx context.Context, req *dto.SignUpRequest) (string, string, error)
	SignOut(ctx context.Context) error
}

type UserService struct {
	validator   validation.Validation
	userRepo    repository.IUserRepository
	minioClient *minio.MinioClient
}

func NewUserService(validator validation.Validation, userRepo repository.IUserRepository, minioClient *minio.MinioClient) *UserService {
	return &UserService{
		validator:   validator,
		userRepo:    userRepo,
		minioClient: minioClient,
	}
}

func (u *UserService) SignIn(ctx context.Context, req *dto.SignInRequest) (string, string, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return "", "", err
	}
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Errorf("Login.GetUserByEmail fail, email: %s, error: %s", req.Email, err)
		return "", "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", "", errors.New("wrong message")
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}

	accessToken := jwt.GenerateAccessToken(tokenData)
	refreshToken := jwt.GenerateRefreshToken(tokenData)

	return accessToken, refreshToken, nil
}

func (u *UserService) SignUp(ctx context.Context, req *dto.SignUpRequest) (string, string, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return "", "", err
	}

	var avatarUrlUpload = ""
	if req.Avatar != nil {
		//logger.Fatal(req.Avatar)
		avatarURL, err := u.minioClient.UploadFile(ctx, req.Avatar, "users")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return "", "", err
		}
		avatarUrlUpload = avatarURL
	}

	var user *model.User
	utils.MapStruct(&user, &req)
	user.AvatarUrl = avatarUrlUpload

	err := u.userRepo.Create(ctx, user)
	if err != nil {
		logger.Errorf("Register.Create fail, email: %s, error: %s", req.Email, err)
		return "", "", err
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}

	accessToken := jwt.GenerateAccessToken(tokenData)
	refreshToken := jwt.GenerateRefreshToken(tokenData)

	return accessToken, refreshToken, nil
}

func (u *UserService) SignOut(ctx context.Context) error {
	return nil
}
