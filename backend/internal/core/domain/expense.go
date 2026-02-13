package domain

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	GroupID     uuid.UUID      `json:"group_id"`
	PayerID     uuid.UUID      `json:"payer_id"`
	Amount      float64        `json:"amount"`
	Description string         `json:"description"`
	Method      string         `json:"method"` // equal, percentage, fixed
	Date        time.Time      `json:"date"`
	Splits      []ExpenseSplit `json:"splits" gorm:"foreignKey:ExpenseID"`
	CreatedAt   time.Time      `json:"created_at"`
}

type ExpenseSplit struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	ExpenseID  uuid.UUID `json:"expense_id"`
	UserID     uuid.UUID `json:"user_id"`
	AmountOwed float64   `json:"amount_owed"`
}
