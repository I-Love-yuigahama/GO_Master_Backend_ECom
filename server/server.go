package server

import (
	"net/http"

	"Ecom.com/TUT/internal/config"
	"Ecom.com/TUT/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	config         *config.Config
	db             *gorm.DB
	logger         *zerolog.Logger
	authServices   *services.AuthServices
	productService *services.ProductServices
	userService    *services.UserService
	uploadService  *services.UploadService
}

func New(cfg *config.Config,
	db *gorm.DB,
	logger *zerolog.Logger,
	authServices *services.AuthServices,
	productService *services.ProductServices,
	userService *services.UserService,
	uploadService *services.UploadService,

) *Server {
	return &Server{
		config:         cfg,
		db:             db,
		logger:         logger,
		authServices:   authServices,
		productService: productService,
		userService:    userService,
		uploadService:  uploadService,
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	// Add midleware

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	// add routes
	router.GET("/health", s.healthCheck)

	router.Static("/uploads", "./uploads")

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			//nolint:gocritic // i need this for readability
			auth.POST("/register", s.register)
			auth.POST("/login", s.login)
			auth.POST("/refresh", s.refreshToken)
			auth.POST("/logout", s.logout)
		}

		protected := api.Group("/")
		protected.Use(s.authMiddleWares())
		{
			// user route
			users := protected.Group("/users")
			{
				userRoutes := users
				userRoutes.GET("/profile", s.getProfile)
				userRoutes.PUT("/profile", s.updateProfile)
			}

			// category

			categories := protected.Group("/categories")
			{
				categyRoute := categories
				categyRoute.POST("/", s.adminMiddleWares(), s.createCatogry)
				categyRoute.PUT("/:id", s.adminMiddleWares(), s.updateCategory)
				categyRoute.DELETE("/:id", s.adminMiddleWares(), s.deleteCategory)
			}

			// product

			products := protected.Group("/products")
			{
				productRoute := products
				productRoute.POST("/", s.adminMiddleWares(), s.createProduct)
				productRoute.PUT("/:id", s.adminMiddleWares(), s.updateProduct)
				productRoute.DELETE("/:id", s.adminMiddleWares(), s.deleteProduct)
				productRoute.POST("/:id/images", s.adminMiddleWares(), s.uploadProductImage)

			}

		}

		// public route

		api.GET("/categories", s.getCategories)
		api.GET("/products", s.getProducts)
		api.GET("/products/:id", s.getProduct)

	}

	return router
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Status": "ok"})
}

func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
