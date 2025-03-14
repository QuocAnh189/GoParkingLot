package service

import (
	"context"
	"errors"
	"fmt"
	"goparking/domains/auth/dto"
	"goparking/domains/auth/model"
	"goparking/domains/auth/repository"
	"goparking/internals/libs/logger"
	"goparking/internals/libs/validation"
	"goparking/pkgs/minio"
	"goparking/pkgs/redis"
	"goparking/pkgs/token"
	"goparking/pkgs/utils"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	SignIn(ctx context.Context, req *dto.SignInRequest) (string, string, *model.User, error)
	SignUp(ctx context.Context, req *dto.SignUpRequest) (string, string, *model.User, error)
	DeleteUser(ctx context.Context, id string) error
	SignOut(ctx context.Context, userID string, jit string) error
}

type UserService struct {
	validator   validation.Validation
	userRepo    repository.IUserRepository
	minioClient *minio.MinioClient
	cache       redis.IRedis
	token       token.IMarker
}

func NewUserService(
	validator validation.Validation,
	userRepo repository.IUserRepository,
	minioClient *minio.MinioClient,
	cache redis.IRedis,
	token token.IMarker,
) *UserService {
	return &UserService{
		validator:   validator,
		userRepo:    userRepo,
		minioClient: minioClient,
		cache:       cache,
		token:       token,
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

	tokenData := token.AuthPayload{
		ID:    user.ID,
		Email: user.Email,
		Role:  "",
	}

	accessToken := u.token.GenerateAccessToken(&tokenData)
	refreshToken := u.token.GenerateRefreshToken(&tokenData)

	return accessToken, refreshToken, user, nil
}

func (u *UserService) SignUp(ctx context.Context, req *dto.SignUpRequest) (string, string, *model.User, error) {
	if err := u.validator.ValidateStruct(req); err != nil {
		return "", "", nil, err
	}

	var avatarUrlUpload = ""
	if req.Avatar != nil {
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

	tokenData := token.AuthPayload{
		ID:    user.ID,
		Email: user.Email,
		Role:  "",
	}

	accessToken := u.token.GenerateAccessToken(&tokenData)
	refreshToken := u.token.GenerateRefreshToken(&tokenData)

	return accessToken, refreshToken, user, nil
}

func (u *UserService) DeleteUser(ctx context.Context, id string) error {
	user, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		return err
	}

	if err := u.userRepo.Delete(ctx, user); err != nil {
		return err
	}

	u.minioClient.DeleteFile(ctx, user.AvatarUrl)

	return nil
}

func (u *UserService) SignOut(ctx context.Context, userID string, jit string) error {
	value := `{"status": "blacklisted"}`

	// err := u.cache.Set(fmt.Sprintf("blacklist:%s", strings.ReplaceAll(token, " ", "_")), value)
	err := u.cache.Set(fmt.Sprintf("blacklist:%s_%s", userID, jit), value)
	if err != nil {
		logger.Error("Failed to blacklist token: ", err)
		return err
	}

	logger.Info("User signed out successfully")
	return nil
}
