package services

import (
	"Ecom.com/TUT/dto"
	"Ecom.com/TUT/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService{
	return &UserService{
		db: db,
	}
}

func(s *UserService) GetProfile(userId uint )(*dto.UserResponse, error){

	var user models.User

	if err := s.db.First(&user,userId).Error; err!= nil{
		return nil,err
	}

	return &dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
			Role:      string(user.Role),
			IsActive:  user.IsActive,
	},nil



}

func(s *UserService)UpdateProfile(userId uint,req *dto.UpdateProfileRequest, )(*dto.UserResponse,error){

	var user models.User

	if err := s.db.First(&user,userId).Error; err!= nil{
		return nil,err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Phone = req.Phone

	if err := s.db.Save(&user).Error; err != nil {
		return nil,err
	}

	return s.GetProfile(userId)


}