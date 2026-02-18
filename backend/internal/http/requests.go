package http

import "github.com/google/uuid"

type registerRequest struct {
	Name     string `json:"name" binding:"required" example:"Teste teste"`
	Email    string `json:"email" binding:"required" example:"teste@teste.com"`
	Password string `json:"password" binding:"required" example:"senha12345"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required, email" example:"teste@teste.com"`
	Password string `json:"password" binding:"required" example:"senha12345"`
}

type CreateExpenseRequest struct {
	GroupID uuid.UUID      `json:"group_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title   string         `json:"title" binding:"required" example:"Almoço"`
	Amount  float64        `json:"amount" binding:"required,gt=0" example:"100.50"`
	PayerID uuid.UUID      `json:"payer_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Splits  []SplitRequest `json:"splits" binding:"required,gt=0"`
}

type SplitRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	Amount float64   `json:"amount" binding:"required,gt=0" example:"50.25"`
}

type createGroupRequest struct {
	Title       string `json:"title" binding:"required" example:"Grupo de Viagem"`
	Description string `json:"description" example:"Grupo para organizar uma viagem"`
}
