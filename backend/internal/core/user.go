package user

import "github.com/pairbytes-dev/tripsplit-monorepo/tree/dev/backend/internal/security"

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	IsActive     bool
}

func NewUser(id int64, name, email, rawPassword string) (*User, error) {
	hashed, err := security.HashPassword(rawPassword)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           id,
		Name:         name,
		Email:        email,
		PasswordHash: hashed,
		IsActive:     true,
	}, nil
}
