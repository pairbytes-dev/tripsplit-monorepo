package db

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/core/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) error {
	m := domain.ToModel(u)

	result := r.db.WithContext(ctx).Create(m)

	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == "23505" {
				log.Printf("[QA ALERT] Tentativa de cadastro com email duplicado: %s", u.Email)
				return errors.New("este email já está cadastrado no sistema")
			}
		}

		log.Printf("[ERROR] Falha técnica ao criar usuário: %v", result.Error)
		return result.Error
	}

	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var m domain.UserModel
	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return domain.ToDomain(&m), nil
}
