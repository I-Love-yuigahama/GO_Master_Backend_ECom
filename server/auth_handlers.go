package server

import (
	"Ecom.com/TUT/dto"
	"Ecom.com/TUT/utils"
	"github.com/gin-gonic/gin"
)

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

func (s *Server) getProfile(ctx *gin.Context) {

	userId := ctx.GetUint("user_id")
	profile, err := s.userService.GetProfile(userId)

	if err != nil {
		utils.NotFoundResponse(ctx, "Profile not found ")
		return
	}

	utils.SuccessResponse(ctx, "Profile retrived successfully", profile)
}

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
