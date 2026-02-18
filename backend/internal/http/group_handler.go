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

// CreateGroup godoc
// @Summary      Criar novo grupo
// @Description  Cria um novo grupo de despesas, associando-o ao usuário autenticado como proprietário.
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        group body createGroupRequest true "Dados do grupo"
// @Success      201   {object}  domain.Group
// @Failure      400   {object}  map[string]string
// @Failure		 401   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /groups [post]
func (h *GroupHandler) CreateGroup(c *gin.Context) {
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

	if err := h.repo.CreateGroup(c.Request.Context(), group, ownerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao criar grupo"})
		return
	}

	c.JSON(http.StatusCreated, group)
}
