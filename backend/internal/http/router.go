package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/db"
	"gorm.io/gorm"
)

func NewRouter(gormDB *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := db.NewUserRepository(gormDB)
	authHandler := NewAuthHandler(userRepo)

	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", authHandler.Register)
			authGroup.POST("/login", authHandler.Login)
		}
	}

	return r
}
