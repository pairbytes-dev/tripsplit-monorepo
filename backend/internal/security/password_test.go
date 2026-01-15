package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_WithValidPassword_ReturnHashedPassword(t *testing.T) {
	//arrange
	plain := "Senha1234"

	//act
	hash, err := HashPassword(plain)

	//assert
	assert.NoError(t, err, "HashPassword deve retornar nil error para senha válida")
	assert.NotEmpty(t, hash, "HashPassword deve retornar hash não vazio")
	assert.NotEqual(t, plain, hash, "Hash não deve ser igual à senha em texto puro")

	ok := CheckPassword(plain, hash)
	assert.True(t, ok, "CheckPassword deve retornar true para senha correta")
}
