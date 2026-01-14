package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/core/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	query := `
	INSERT INTO users (name, email, password_hash, is_active)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return r.db.QueryRowContext(
		ctx,
		query,
		u.Name,
		u.Email,
		u.PasswordHash,
		u.IsActive,
	).Scan(&u.ID)
}
