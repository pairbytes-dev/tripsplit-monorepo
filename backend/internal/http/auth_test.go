package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/core/user"
	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/security"
)

// fake de UserRepository para os testes
type fakeUserRepo struct {
	// configuráveis
	createErr     error
	getByEmailErr error
	userToReturn  *user.User

	// para inspeção
	createdUser *user.User
	getEmailArg string
}

func (f *fakeUserRepo) Create(ctx context.Context, u *user.User) error {
	f.createdUser = u
	return f.createErr
}

func (f *fakeUserRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	f.getEmailArg = email
	if f.getByEmailErr != nil {
		return nil, f.getByEmailErr
	}
	return f.userToReturn, nil
}

func init() {
	gin.SetMode(gin.TestMode)
}

// helper para gerar hash de senha nos testes
func mustHash(t *testing.T, plain string) string {
	t.Helper()
	hash, err := security.HashPassword(plain)
	if err != nil {
		t.Fatalf("erro gerando hash de senha no teste: %v", err)
	}
	return hash
}

//
// ---------- TESTES DE REGISTER ----------
//

func TestAuthHandler_Register_WithValidPayload_Returns201AndCreatesUser(t *testing.T) {
	// Arrange
	repo := &fakeUserRepo{}
	h := NewAuthHandler(repo)

	body := map[string]string{
		"name":     "Andre",
		"email":    "Andre@example.com",
		"password": "Senha1234",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodPost, "/v1/auth/register", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	r.POST("/v1/auth/register", h.Register)

	// Act
	r.ServeHTTP(w, c.Request)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var resp map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "Usuário criado com sucesso", resp["message"])

	userMap, ok := resp["user"].(map[string]any)
	assert.True(t, ok)

	assert.Equal(t, "Andre", userMap["name"])
	assert.Equal(t, "andre@example.com", userMap["email"])

	assert.NotNil(t, repo.createdUser)
	assert.Equal(t, "Andre", repo.createdUser.Name)
	assert.Equal(t, "andre@example.com", repo.createdUser.Email)
}

func TestAuthHandler_Register_WithInvalidPayload_Returns400(t *testing.T) {
	// Arrange
	repo := &fakeUserRepo{}
	h := NewAuthHandler(repo)

	// payload inválido: faltando campos obrigatórios
	body := map[string]string{
		"email": "sem_nome@example.com",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodPost, "/v1/auth/register", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	r.POST("/v1/auth/register", h.Register)

	// Act
	r.ServeHTTP(w, c.Request)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid payload", resp["error"])
	assert.Nil(t, repo.createdUser)
}

func TestAuthHandler_Register_WhenCreateFails_Returns500(t *testing.T) {
	// Arrange
	repo := &fakeUserRepo{
		createErr: errors.New("db error"),
	}
	h := NewAuthHandler(repo)

	body := map[string]string{
		"name":     "Andre",
		"email":    "andre@example.com",
		"password": "Senha1234",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodPost, "/v1/auth/register", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	r.POST("/v1/auth/register", h.Register)

	// Act
	r.ServeHTTP(w, c.Request)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var resp map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "could not create user", resp["error"])
	assert.NotNil(t, repo.createdUser)
}

//
// ---------- TESTES DE LOGIN ----------
//

func TestAuthHandler_Login_WithInvalidPayload_Returns400(t *testing.T) {
	// Arrange
	repo := &fakeUserRepo{}
	h := NewAuthHandler(repo)

	// payload inválido: sem email
	body := map[string]string{
		"password": "Senha1234",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodPost, "/v1/auth/login", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	r.POST("/v1/auth/login", h.Login)

	// Act
	r.ServeHTTP(w, c.Request)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid payload", resp["error"])
}

func TestAuthHandler_Login_WhenUserNotFound_Returns401(t *testing.T) {
	// Arrange
	repo := &fakeUserRepo{
		getByEmailErr: errors.New("not found"),
	}
	h := NewAuthHandler(repo)

	body := map[string]string{
		"email":    "naoexiste@example.com",
		"password": "Senha1234",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodPost, "/v1/auth/login", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	r.POST("/v1/auth/login", h.Login)

	// Act
	r.ServeHTTP(w, c.Request)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var resp map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "Email or password is incorrect", resp["error"])
}

func TestAuthHandler_Login_WithWrongPassword_Returns401(t *testing.T) {
	// Arrange
	hashed := mustHash(t, "SenhaCorreta")

	repo := &fakeUserRepo{
		userToReturn: &user.User{
			ID:           1,
			Name:         "Andre",
			Email:        "andre@example.com",
			PasswordHash: hashed,
		},
	}
	h := NewAuthHandler(repo)

	body := map[string]string{
		"email":    "andre@example.com",
		"password": "SenhaErrada",
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodPost, "/v1/auth/login", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	r.POST("/v1/auth/login", h.Login)

	// Act
	r.ServeHTTP(w, c.Request)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var resp map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "Email or password is incorrect", resp["error"])
}

func TestAuthHandler_Login_WithValidCredentials_Returns200AndToken(t *testing.T) {
	// Arrange
	plainPassword := "Senha1234"
	hashed := mustHash(t, plainPassword)

	repo := &fakeUserRepo{
		userToReturn: &user.User{
			ID:           1,
			Name:         "Andre",
			Email:        "andre@example.com",
			PasswordHash: hashed,
		},
	}
	h := NewAuthHandler(repo)

	body := map[string]string{
		"email":    "andre@example.com",
		"password": plainPassword,
	}
	b, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodPost, "/v1/auth/login", bytes.NewReader(b))
	c.Request.Header.Set("Content-Type", "application/json")

	r.POST("/v1/auth/login", h.Login)

	// Act
	r.ServeHTTP(w, c.Request)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]any
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "Login efetuado com sucesso", resp["message"])

	token, ok := resp["token"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, token)
}
