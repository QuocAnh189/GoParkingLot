package http

import (
	"github.com/gin-gonic/gin"
	"goparking/domains/auth/dto"
	"goparking/domains/auth/service"
	"goparking/internals/libs/logger"
	"goparking/pkgs/response"
	"net/http"
)

type AuthHandler struct {
	service service.IUserService
}

func NewAuthHandler(service service.IUserService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	accessToken, refreshToken, err := h.service.SignIn(c, &req)

	if err != nil {
		logger.Error("Failed to sign up ", err)
		switch err.Error() {
		case "wrong message":
			response.Error(c, http.StatusConflict, err, "wrong message")
			return
		}

		response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		return
	}

	var res dto.SignInResponse
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	response.JSON(c, http.StatusOK, res)
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	accessToken, refreshToken, err := h.service.SignUp(c, &req)
	if err != nil {
		logger.Error("Failed to sign up ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		return
	}

	var res dto.SignUpResponse
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	response.JSON(c, http.StatusOK, res)
}

func (h *AuthHandler) SignOut(c *gin.Context) {
	response.JSON(c, 200, "SignOut")
}
