package server

import (
	"Ecom.com/TUT/dto"
	"Ecom.com/TUT/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Register a new user
// @Description Create a new user account with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "User registration data"
// @Success 201 {object} utils.Response{data=dto.AuthResponse} "User registered successfully"
// @Failure 400 {object} utils.Response "Invalid request data or user already exists"
// @Router /auth/register [post]

func (s *Server) register(ctx *gin.Context) {

	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data ", err)
		return
	}

	response, err := s.authServices.Register(&req)

	if err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data ", err)
		return
	}

	utils.CreatedResponse(ctx, "User registered successfully", response)

}

func (s *Server) login(ctx *gin.Context) {

	var req dto.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data ", err)
		return
	}

	response, err := s.authServices.Login(&req)

	if err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data ", err)
		return
	}

	utils.SuccessResponse(ctx, "Login successfully", response)

}

// @Summary Refresh access token
// @Description Get a new access token using refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} utils.Response{data=dto.AuthResponse} "Token refreshed successfully"
// @Failure 401 {object} utils.Response "Invalid refresh token"
// @Router /auth/refresh [post]

func (s *Server) refreshToken(ctx *gin.Context) {

	var req dto.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data ", err)
		return
	}

	response, err := s.authServices.RefreshToken(&req)

	if err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data ", err)
		return
	}

	utils.SuccessResponse(ctx, "Login successfully", response)

}

func (s *Server) logout(ctx *gin.Context) {

	var req dto.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data ", err)
		return
	}

	if err := s.authServices.Logout(req.RefreshToken); err != nil {
		utils.InternalServerErrorResponse(ctx, "Logout Failed", err)
	}

	utils.SuccessResponse(ctx, "Login successfully", nil)

}


// @Summary Update user profile
// @Description Update current authenticated user's profile information
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateProfileRequest true "Profile update data"
// @Success 200 {object} utils.Response{data=dto.UserResponse} "Profile updated successfully"
// @Failure 400 {object} utils.Response "Invalid request data"
// @Failure 401 {object} utils.Response "Unauthorized"
// @Router /users/profile [put]
func (s *Server) getProfile(ctx *gin.Context) {

	userId := ctx.GetUint("user_id")
	profile, err := s.userService.GetProfile(userId)

	if err != nil {
		utils.NotFoundResponse(ctx, "Profile not found ")
		return
	}

	utils.SuccessResponse(ctx, "Profile retrived successfully", profile)
}


// @Summary Update user profile
// @Description Update current authenticated user's profile information
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateProfileRequest true "Profile update data"
// @Success 200 {object} utils.Response{data=dto.UserResponse} "Profile updated successfully"
// @Failure 400 {object} utils.Response "Invalid request data"
// @Failure 401 {object} utils.Response "Unauthorized"
// @Router /users/profile [put]
func (s *Server) updateProfile(ctx *gin.Context) {

	userId := ctx.GetUint("user_id")

	var req dto.UpdateProfileRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid requiest data ", err)
		return
	}

	updatedProfile, err := s.userService.UpdateProfile(userId, &req)

	if err != nil {
		utils.BadRequestResponse(ctx, "cant update the Profile", err)
		return
	}

	utils.SuccessResponse(ctx, "Updated the Profile ", updatedProfile)

}
