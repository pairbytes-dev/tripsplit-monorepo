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

// CreateExpense godoc
// @Summary      Criar nova despesa
// @Description  Registra uma nova despesa associada a um grupo, com detalhes de quem pagou e como a despesa deve ser dividida entre os membros.
// @Tags         expenses
// @Accept       json
// @Produce      json
// @Param        expense body CreateExpenseRequest true "Dados da despesa"
// @Success      201   {object}  domain.Expense
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /expenses [post]
func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
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

	if err := h.repo.CreateExpense(c.Request.Context(), expense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao registrar despesa no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, expense)
}
