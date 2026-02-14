package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/core/domain"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) Create(ctx context.Context, g *domain.Group, ownerID uuid.UUID) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(g).Error; err != nil {
			return err
		}
		member := &domain.GroupMember{
			GroupID: g.ID,
			UserID:  ownerID,
			Role:    "owner",
		}
		if err := tx.Create(member).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *GroupRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Group, error) {
	var groups []domain.Group
	err := r.db.WithContext(ctx).
		Joins("JOIN group_members ON group_members.group_id = groups.id").
		Where("group_members.user_id = ?", userID).
		Find(&groups).Error

	return groups, err
}
