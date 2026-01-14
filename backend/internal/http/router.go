package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/db"
)

func NewRouter(conn *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepo := db.NewUserRepository(conn)
	authHandler := NewAuthHandler(userRepo)

	v1 := r.Group("/V1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
		}
	}

	return r

}
