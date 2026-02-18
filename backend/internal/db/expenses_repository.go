package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/core/domain"
	"gorm.io/gorm"
)

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) CreateExpense(ctx context.Context, expense *domain.Expense) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(expense).Error; err != nil {
			return err
		}

		if len(expense.Splits) > 0 {
			if err := tx.Create(&expense.Splits).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *ExpenseRepository) GetByGroupID(ctx context.Context, groupID uuid.UUID) ([]domain.Expense, error) {
	var expenses []domain.Expense
	err := r.db.WithContext(ctx).
		Preload("Splits").
		Where("group_id = ?", groupID).
		Find(&expenses).Error

	return expenses, err
}
