package server

import (
	"Ecom.com/TUT/dto"
	"Ecom.com/TUT/internal/services"
	"Ecom.com/TUT/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) register(ctx *gin.Context) {

	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data ", err)
		return
	}

	authService := services.NewAuthService(s.db, s.config)
	response, err := authService.Register(&req)

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

	authservices := services.NewAuthService(s.db, s.config)
	response, err := authservices.Login(&req)

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

	authservices := services.NewAuthService(s.db, s.config)
	response, err := authservices.RefreshToken(&req)

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

	authservices := services.NewAuthService(s.db, s.config)
	if err := authservices.Logout(req.RefreshToken); err != nil {
		utils.InternalServerErrorResponse(ctx, "Logout Failed", err)
	}

	utils.SuccessResponse(ctx, "Login successfully", nil)

}

func (s *Server) getProfile(ctx *gin.Context){

	userId := ctx.GetUint("user_id")
	userService := services.NewUserService(s.db)
	profile , err := userService.GetProfile(userId)

	if err != nil{
		utils.NotFoundResponse(ctx,"Profile not found ")
		return
	}

	utils.SuccessResponse(ctx,"Profile retrived successfully",profile)
}

func (s *Server) updateProfile(ctx *gin.Context){

	userId := ctx.GetUint("user_id")

	var req dto.UpdateProfileRequest

	if err := ctx.ShouldBindJSON(&req); err != nil{
		utils.BadRequestResponse(ctx,"Invalid requiest data ",err)
		return
	}

	userServices := services.NewUserService(s.db)
	updatedProfile, err := userServices.UpdateProfile(userId,&req)

	if err !=nil{
		utils.BadRequestResponse(ctx,"cant update the Profile",err)
		return
	}

	utils.SuccessResponse(ctx,"Updated the Profile ",updatedProfile)

}