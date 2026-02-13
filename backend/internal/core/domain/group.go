package domain

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

type Group struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title       string    `gorm:"not null"`
	Description string
	Members     []GroupMember `gorm:"foreignKey:GroupID"`
	Expenses    []Expense     `gorm:"foreignKey:GroupID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GroupMember struct {
	GroupID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Role     string    `gorm:"type:group_role;default:member"` // owner ou member
	JoinedAt time.Time
}
