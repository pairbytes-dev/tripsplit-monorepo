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

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var m user.UserModel
	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&m).Error; err != nil {
		return nil, err
	}
	return user.ToDomain(&m), nil
}
