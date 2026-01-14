package db

import (
	"context"

	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/core/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	m := user.ToModel(u)
	return r.db.WithContext(ctx).Create(m).Error
}
