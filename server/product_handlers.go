package server

import (
	"strconv"

	"Ecom.com/TUT/dto"
	"Ecom.com/TUT/internal/services"
	"Ecom.com/TUT/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) createCatogry(ctx *gin.Context) {
	var req dto.CreateCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid data", err)
		return
	}

	productService := services.NewProductServices(s.db)
	category, err := productService.CreateCatogories(&req)
	if err != nil {
		utils.InternalServerErrorResponse(ctx, "Failed to create category", err)
		return
	}

	utils.CreatedResponse(ctx, "Category created successfully", category)
}

func (s *Server) getCategories(c *gin.Context) {
	productService := services.NewProductServices(s.db)
	categories, err := productService.GetCatogories()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to fetch categories", err)
		return
	}

	utils.SuccessResponse(c, "Categories retrieved successfully", categories)
}

func (s *Server) updateCategory(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid category ID", err)
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	productService := services.NewProductServices(s.db)
	category, err := productService.UpdateCatogory(uint(id), &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update category", err)
		return
	}

	utils.SuccessResponse(c, "Category updated successfully", category)
}

func (s *Server) deleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid category ID", err)
		return
	}

	productService := services.NewProductServices(s.db)
	if err := productService.DeleteCategory(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to delete category", err)
		return
	}

	utils.SuccessResponse(c, "Category deleted successfully", nil)
}