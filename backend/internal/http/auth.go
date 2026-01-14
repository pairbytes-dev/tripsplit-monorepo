package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/core/user"
	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/db"
)

type AuthHandler struct {
	users *db.UserRepository
}

func NewAuthHandler(users *db.UserRepository) *AuthHandler {
	return &AuthHandler{users: users}
}

type registerRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" bindig:"required,email"`
	Password string `jsonn:"password" binding:"required,min=8"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	u, err := user.NewUser(0, req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.users.Create(c.Request.Context(), u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
	})

}
