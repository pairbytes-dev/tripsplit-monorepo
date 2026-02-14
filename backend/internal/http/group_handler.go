package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/core/domain"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/db"
)

type GroupHandler struct {
	repo *db.GroupRepository
}

func NewGroupHandler(repo *db.GroupRepository) *GroupHandler {
	return &GroupHandler{repo: repo}
}

type createGroupRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

func (h *GroupHandler) Create(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "usuário não autenticado"})
		return
	}
	ownerID := userIDVal.(uuid.UUID)

	var req createGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos: título é obrigatório"})
		return
	}

	group := &domain.Group{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
	}

	if err := h.repo.Create(c.Request.Context(), group, ownerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao criar grupo"})
		return
	}

	c.JSON(http.StatusCreated, group)
}
