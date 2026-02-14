package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/core/domain"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/db"
)

type ExpenseHandler struct {
	repo *db.ExpenseRepository
}

func NewExpenseHandler(repo *db.ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: repo}
}

type CreateExpenseRequest struct {
	GroupID uuid.UUID      `json:"group_id" binding:"required"`
	Title   string         `json:"title" binding:"required"`
	Amount  float64        `json:"amount" binding:"required,gt=0"`
	PayerID uuid.UUID      `json:"payer_id" binding:"required"`
	Splits  []SplitRequest `json:"splits" binding:"required,gt=0"`
}

type SplitRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
	Amount float64   `json:"amount" binding:"required,gt=0"`
}

func (h *ExpenseHandler) Create(c *gin.Context) {
	var req CreateExpenseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos ou campos obrigatórios ausentes"})
		return
	}

	expenseID := uuid.New()
	expense := &domain.Expense{
		ID:          expenseID,
		GroupID:     req.GroupID,
		Description: req.Title,
		Amount:      req.Amount,
		PayerID:     req.PayerID,
	}

	var splits []domain.ExpenseSplit
	for _, s := range req.Splits {
		splits = append(splits, domain.ExpenseSplit{
			ID:         uuid.New(),
			ExpenseID:  expenseID,
			UserID:     s.UserID,
			AmountOwed: s.Amount,
		})
	}

	expense.Splits = splits

	if err := h.repo.Create(c.Request.Context(), expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao registrar despesa no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, expense)
}
