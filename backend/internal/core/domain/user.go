package domain

import (
	"errors"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/security"
)

var (
	ErrInvalidName  = errors.New("O nome é necessário")
	ErrInvalidEmail = errors.New("O email é necessário")
	ErrWeakPassword = errors.New("A senha precisa ter pelo menons 8 caracteres e conter letras e números")
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	IsActive     bool      `json:"is_active"`
}

func NewUser(name, email, rawPassword string) (*User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(strings.ToLower(email))

	if name == "" {
		return nil, ErrInvalidName
	}
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if rawPassword == "" || len(rawPassword) < 8 {
		return nil, ErrWeakPassword
	}
	if !isStrongPassword(rawPassword) {
		return nil, ErrWeakPassword
	}

	hashed, err := security.HashPassword(rawPassword)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: hashed,
		IsActive:     true,
	}, nil
}

func isStrongPassword(p string) bool {
	if len(p) < 8 {
		return false
	}

	hasLetter := false
	hasDigit := false

	for _, r := range p {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}
	return hasLetter && hasDigit
}

type UserModel struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name         string    `gorm:"not null"`
	Email        string    `gorm:"not null;"`
	PasswordHash string    `gorm:"not null"`
	IsActive     bool      `gorm:"not null;default:true"`
}

func (UserModel) TableName() string {
	return "users"
}

func ToModel(u *User) *UserModel {
	if u == nil {
		return nil
	}
	return &UserModel{
		ID:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		IsActive:     u.IsActive,
	}
}

func ToDomain(m *UserModel) *User {
	if m == nil {
		return nil
	}
	return &User{
		ID:           m.ID,
		Name:         m.Name,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		IsActive:     m.IsActive,
	}
}
