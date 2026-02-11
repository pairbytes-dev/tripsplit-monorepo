package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pairbytes-dev/tripsplit-monorepo/internal/core/user"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/db"
	httpapi "github.com/pairbytes-dev/tripsplit-monorepo/internal/http"
)

func main() {
	// Funções auxiliares para ler variáveis de ambiente com um valor padrão (fallback)
	getEnv := func(key, defaultValue string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultValue
	}

	getEnvInt := func(key string, defaultValue int) int {
		if value := os.Getenv(key); value != "" {
			if i, err := strconv.Atoi(value); err == nil {
				return i
			}
		}
		return defaultValue
	}

	// Agora a configuração prioriza o que vem do Docker Compose
	cfg := db.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvInt("DB_PORT", 5432),
		User:     getEnv("DB_USER", "tripsplit"),
		Password: getEnv("DB_PASSWORD", "tripsplit"),
		DBName:   getEnv("DB_NAME", "tripsplit"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	gormDB, err := db.OpenGormPostgres(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco (%s:%d): %v", cfg.Host, cfg.Port, err)
	}

	if err := gormDB.AutoMigrate(&user.UserModel{}); err != nil {
		log.Fatal("Erro na migração:", err)
	}

	router := httpapi.NewRouter(gormDB)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	fmt.Println("Rotas registradas:")
	for _, route := range router.Routes() {
		fmt.Printf("  %-6s %s\n", route.Method, route.Path)
	}

	fmt.Printf("\nBackend rodando em: http://localhost:8080 (Conectado em: %s)\n", cfg.Host)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
