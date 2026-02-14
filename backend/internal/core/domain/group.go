package domain

import (
	"time"

	uuid "github.com/google/uuid"
)

type Group struct {
	ID          uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title       string        `json:"title" gorm:"not null"`
	Description string        `json:"description"`
	Members     []GroupMember `json:"members" gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Expenses    []Expense     `json:"expenses" gorm:"foreignKey:GroupID"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type GroupMember struct {
	GroupID  uuid.UUID `json:"group_id" gorm:"type:uuid;primaryKey"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey"`
	Role     string    `json:"role" gorm:"default:member"`
	JoinedAt time.Time `json:"joined_at" gorm:"autoCreateTime"`
}
