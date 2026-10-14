package server

import (
	"strings"

	"Ecom.com/TUT/models"
	"Ecom.com/TUT/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) authMiddleWares() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Authorization : Bearer JWT

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.UnauthorizedResponse(ctx, "Authorization Header required")
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.UnauthorizedResponse(ctx, "Invalid Authorization Header required")
			ctx.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenParts[1], s.config.JWT.Secret)
		if err != nil {
			utils.UnauthorizedResponse(ctx, "Invalid Authorization Header required")
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_email", claims.Email)
		ctx.Set("user_role", claims.Role)

		ctx.Next()

	}
}

func (s *Server) adminMiddleWares() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("user_role")

		if !exists {
			utils.UnauthorizedResponse(ctx, "Forbidden")
			ctx.Abort()
			return
		}

		if role != string(models.UserRoleAdmin) {
			utils.UnauthorizedResponse(ctx, "Forbidden")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
