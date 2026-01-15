package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser_WithValidData_CreatesActiveUserWithNormalizedFields(t *testing.T) {
	//Arrange
	inputName := " Andre "
	inputEmail := " Andre@Example.com "
	inputPassword := "Senha1234"

	//Act
	u, err := NewUser(0, inputName, inputEmail, inputPassword)

	//Assert
	assert.NoError(t, err, "NewUser deve retornar nil error com dados válidos")
	assert.NotNil(t, u, "NewUser deve retornar um usuário não nil")

	assert.Equal(t, "Andre", u.Name, "Name deve ser trimado")
	assert.Equal(t, "andre@example.com", u.Email, "Email deve ser trimado e em lowercase")
	assert.True(t, u.IsActive, "Usuário recém-criado deve estar ativo")
	assert.NotEmpty(t, u.PasswordHash, "PasswordHash deve ser gerado")
	assert.NotEqual(t, inputPassword, u.PasswordHash, "Senha NÃO deve ser armazenada em texto puro")
}

func TestNewUser_WithInvalidData_ReturnsExpectedErrors(t *testing.T) {
	type testCase struct {
		name        string
		inputName   string
		inputEmail  string
		inputPass   string
		expectedErr error
	}

	cases := []testCase{
		{
			name:        "nome vazio",
			inputName:   "   ",
			inputEmail:  "andre@example.com",
			inputPass:   "Senha1234",
			expectedErr: ErrInvalidName,
		},
		{
			name:        "email vazio",
			inputName:   "Andre",
			inputEmail:  "   ",
			inputPass:   "Senha1234",
			expectedErr: ErrInvalidEmail,
		},
		{
			name:        "senha curta",
			inputName:   "Andre",
			inputEmail:  "andre@example.com",
			inputPass:   "Abc123",
			expectedErr: ErrWeakPassword,
		},
		{
			name:        "senha só letras",
			inputName:   "Andre",
			inputEmail:  "andre@example.com",
			inputPass:   "abcdefgh",
			expectedErr: ErrWeakPassword,
		},
		{
			name:        "senha só dígitos",
			inputName:   "Andre",
			inputEmail:  "andre@example.com",
			inputPass:   "12345678",
			expectedErr: ErrWeakPassword,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := NewUser(0, tc.inputName, tc.inputEmail, tc.inputPass)

			assert.Nil(t, u, "Para dados inválidos, o usuário retornado deve ser nil")
			assert.Error(t, err, "Para dados inválidos, deve retornar um erro")
			assert.Equal(t, tc.expectedErr, err, "Erro retornado deve ser o esperado")
		})
	}
}
