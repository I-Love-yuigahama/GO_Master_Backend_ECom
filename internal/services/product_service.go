package services

import (
	"Ecom.com/TUT/dto"
	"Ecom.com/TUT/models"
	"gorm.io/gorm"
)

type ProductServices struct {
	db *gorm.DB
}

func NewProductServices(db *gorm.DB) *ProductServices{
	return &ProductServices{db: db}
}

func(s ProductServices) CreateCatofry(req *dto.CreateCategoryRequest)(*dto.CategoryResponse, error){
	
	category := models.Category{
		Name: req.Name,
		Description: req.Description,
	}

	if err := s.db.Create(&category).Error; err!= nil{
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil


}

func(s *ProductServices)GetCatogories()([]dto.CategoryResponse,error){

	var categories []models.Category

	if err:= s.db.Where("is_active = ?",true).Find(&categories).Error; err != nil{
		return nil, err
	}

	response := make([]dto.CategoryResponse,len(categories))

	for i := range categories{
		response[i] = dto.CategoryResponse{
			ID:          categories[i].ID,
			Name:        categories[i].Name,
			Description: categories[i].Description,
			IsActive:    categories[i].IsActive,
		}
	}

	return response,nil
	

}

func(s *ProductServices) UpdateCatogory(id uint, req *dto.UpdateCategoryRequest)(*dto.CategoryResponse, error){

	var category models.Category

	if err:= s.db.First(&category,id).Error; err != nil{
		return nil,err
	}

	category.Name = req.Name
	category.Description = req.Description

	if req.IsActive != nil{
		category.IsActive = *req.IsActive
	}


	if err := s.db.Save(&category).Error;err != nil{
		return nil,err
	}

	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil


}