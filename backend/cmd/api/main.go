package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/pairbytes-dev/tripsplit-monorepo/internal/db"
	httpapi "github.com/pairbytes-dev/tripsplit-monorepo/internal/http"
)

// @title           TripSplit API
// @version         1.0
// @description     API para gestão de despesas em grupo.
// @host            localhost:8080
// @BasePath        /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
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
		Port:     getEnvInt("DB_PORT", 6543),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "postgres"),
		SSLMode:  getEnv("DB_SSLMODE", "require"),
	}

	gormDB, err := db.OpenGormPostgres(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco (%s:%d): %v", cfg.Host, cfg.Port, err)
	}

	fmt.Println("Conexão validada. Schema gerenciado via Migrations Manuais no Supabase.")

	router := httpapi.NewRouter(gormDB)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	fmt.Println("\nRotas registradas:")
	for _, route := range router.Routes() {
		fmt.Printf("  %-6s %s\n", route.Method, route.Path)
	}

	fmt.Printf("\nBackend rodando em: http://localhost:8080 (Conectado em: %s)\n", cfg.Host)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
