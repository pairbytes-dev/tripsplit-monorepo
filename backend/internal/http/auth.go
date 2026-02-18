package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/core/domain"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/db"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/security"
)

// Handler de autenticação
type AuthHandler struct {
	users db.UserRepositoryInterface
}

func NewAuthHandler(users db.UserRepositoryInterface) *AuthHandler {
	return &AuthHandler{users: users}
}

// Register godoc
// @Summary      Cadastrar novo usuário
// @Description  Cria uma nova conta de usuário com nome, e-mail e senha.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      registerRequest  true  "Dados do usuário"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      409   {object}  map[string]string
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	u, err := domain.NewUser(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.users.Create(c.Request.Context(), u); err != nil {
		if err.Error() == "email já cadastrado" {
			c.JSON(http.StatusConflict, gin.H{"error": "Este email já está cadastrado"})
			return
		}
		log.Println("Erro ao criar usuário:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno ao criar usuário"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"user": gin.H{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
		},
	})
}

// Login godoc
// @Summary      Login de usuário
// @Description  Autentica um usuário e retorna um token JWT.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body loginRequest true "Dados de login"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	u, err := h.users.GetByEmail(c.Request.Context(), req.Email)
	if err != nil {
		log.Println("Error finding user:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
		return
	}

	if !security.CheckPassword(req.Password, u.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
		return
	}

	token, err := security.GenerateToken(u.ID, u.Email)
	if err != nil {
		log.Println("Error generating token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login efetuado com sucesso",
		"token":   token,
	})
}
