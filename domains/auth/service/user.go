package service

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"goparking/domains/auth/dto"
	"goparking/domains/auth/model"
	"goparking/domains/auth/repository"
	"goparking/internals/libs/logger"
	"goparking/internals/libs/validation"
	"goparking/pkgs/jwt"
	"goparking/pkgs/minio"
	"goparking/pkgs/redis"
	"goparking/pkgs/utils"
	"strings"
)

type IUserService interface {
	SignIn(ctx context.Context, req *dto.SignInRequest) (string, string, *model.User, error)
	SignUp(ctx context.Context, req *dto.SignUpRequest) (string, string, *model.User, error)
	SignOut(ctx context.Context, userID string, token string) error
}

type UserService struct {
	validator   validation.Validation
	userRepo    repository.IUserRepository
	minioClient *minio.MinioClient
	cache       redis.IRedis
}

func NewUserService(
	validator validation.Validation,
	userRepo repository.IUserRepository,
	minioClient *minio.MinioClient,
	cache redis.IRedis) *UserService {
	return &UserService{
		validator:   validator,
		userRepo:    userRepo,
		minioClient: minioClient,
		cache:       cache,
	}
}

func (u *UserService) SignIn(ctx context.Context, req *dto.SignInRequest) (string, string, *model.User, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return "", "", nil, err
	}
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Errorf("Login.GetUserByEmail fail, email: %s, error: %s", req.Email, err)
		return "", "", nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", "", nil, errors.New("wrong message")
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}

	accessToken := jwt.GenerateAccessToken(tokenData)
	refreshToken := jwt.GenerateRefreshToken(tokenData)

	return accessToken, refreshToken, user, nil
}

func (u *UserService) SignUp(ctx context.Context, req *dto.SignUpRequest) (string, string, *model.User, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return "", "", nil, err
	}

	var avatarUrlUpload = ""
	if req.Avatar != nil {
		//logger.Fatal(req.Avatar)
		avatarURL, err := u.minioClient.UploadFile(ctx, req.Avatar, "users")
		if err != nil {
			logger.Errorf("Failed to upload avatar: %s", err)
			return "", "", nil, err
		}
		avatarUrlUpload = avatarURL
	}

	var user *model.User
	utils.MapStruct(&user, &req)
	user.AvatarUrl = avatarUrlUpload

	err := u.userRepo.Create(ctx, user)
	if err != nil {
		logger.Errorf("Register.Create fail, email: %s, error: %s", req.Email, err)
		return "", "", nil, err
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}

	accessToken := jwt.GenerateAccessToken(tokenData)
	refreshToken := jwt.GenerateRefreshToken(tokenData)

	return accessToken, refreshToken, user, nil
}

func (u *UserService) SignOut(ctx context.Context, userID string, token string) error {
	value := `{"status": "blacklisted"}`

	// Lưu vào Redis với TTL 24 giờ
	err := u.cache.Set(fmt.Sprintf("blacklist:%s", strings.ReplaceAll(token, " ", "_")), value)
	if err != nil {
		logger.Error("Failed to blacklist token: ", err)
		return err
	}

	logger.Info("User signed out successfully")
	return nil
}
