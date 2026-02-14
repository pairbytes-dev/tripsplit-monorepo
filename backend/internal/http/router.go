package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/db"
	"gorm.io/gorm"
)

func NewRouter(gormDB *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userRepo := db.NewUserRepository(gormDB)
	groupRepo := db.NewGroupRepository(gormDB)
	expenseRepo := db.NewExpenseRepository(gormDB)

	authHandler := NewAuthHandler(userRepo)
	groupHandler := NewGroupHandler(groupRepo)
	expenseHandler := NewExpenseHandler(expenseRepo)

	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", authHandler.Register)
			authGroup.POST("/login", authHandler.Login)
		}

		protected := v1.Group("/")
		protected.Use(AuthMiddleware())
		{
			protected.POST("/groups", groupHandler.Create)
			protected.POST("/expenses", expenseHandler.Create)
		}
	}

	return r
}
