package http

import (
	"goparking/domains/auth/dto"
	"goparking/domains/auth/service"
	"goparking/internals/libs/logger"
	"goparking/pkgs/response"
	"goparking/pkgs/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.IUserService
}

func NewAuthHandler(service service.IUserService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

//		@Summary	 Signin a user
//	 @Description Authenticates the user based on the provided credentials and returns a sign-in response if successful.
//		@Tags		 Auth
//		@Produce	 json
//		@Param		 _	body	dto.SignInRequest	  true	"Body"
//		@Success	 200	{object}	response.Response	"Successfully signed in"
//		@Failure	 401	{object}	response.Response	"Unauthorized - Invalid credentials"
//		@Failure	 404	{object}	response.Response	"Not Found - User not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/signin [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	accessToken, refreshToken, user, err := h.service.SignIn(c, &req)

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
	utils.MapStruct(&res.User, user)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Signup a new user
//	 @Description Registers a new user with the provided details. Returns a sign-in response upon successful registration.
//		@Tags		 Auth
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User successfully registered"
//		@Failure	 400	{object}	response.Response	"Invalid user input"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	accessToken, refreshToken, user, err := h.service.SignUp(c, &req)
	if err != nil {
		logger.Error("Failed to sign up ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		return
	}

	var res dto.SignUpResponse
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	utils.MapStruct(&res.User, user)

	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Delete user
//	 @Description Registers a new user with the provided details. Returns a sign-in response upon successful registration.
//		@Tags		 Auth
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User successfully registered"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/delete-user [delete]
func (h *AuthHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	err := h.service.DeleteUser(c, userId)
	if err != nil {
		logger.Error("Failed to delete ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		return
	}

	response.JSON(c, http.StatusOK, true)
}

//		@Summary	 Signout a user
//	 @Description Registers a new user with the provided details. Returns a sign-in response upon successful registration.
//		@Tags		 Auth
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User successfully logout"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/signout [post]
func (h *AuthHandler) SignOut(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		response.Error(c, http.StatusBadRequest, nil, "Missing Authorization header")
		return
	}

	userID, exists := c.Get("userId")
	if !exists {
		response.Error(c, http.StatusNotFound, nil, "Unauthorized")
		return
	}

	err := h.service.SignOut(c, userID.(string), token)
	if err != nil {
		logger.Error("Failed to sign out", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to sign out")
		return
	}

	response.JSON(c, http.StatusOK, "Logout successfully")
}
