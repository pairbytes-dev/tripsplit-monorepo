package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/core/user"
	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/db"
)

// Handler de autenticação
type AuthHandler struct {
	users *db.UserRepository
}

func NewAuthHandler(users *db.UserRepository) *AuthHandler {
	return &AuthHandler{users: users}
}

// Payload de registro
type registerRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
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
		log.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"user": gin.H{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
		},
	})
}
