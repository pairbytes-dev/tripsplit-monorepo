package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/pairbytes-dev/tripsplit-monorepo/internal/core/user"
	"github.com/pairbytes-dev/tripsplit-monorepo/internal/db"
	httpapi "github.com/pairbytes-dev/tripsplit-monorepo/internal/http"
)

func main() {
	cfg := db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "tripsplit",
		Password: "tripsplit",
		DBName:   "tripsplit",
		SSLMode:  "disable",
	}

	gormDB, err := db.OpenGormPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := gormDB.AutoMigrate(&user.UserModel{}); err != nil {
		log.Fatal(err)
	}

	router := httpapi.NewRouter(gormDB)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	fmt.Println("Rotas registradas:")
	for _, route := range router.Routes() {
		fmt.Printf("  %-6s %s\n", route.Method, route.Path)
	}

	fmt.Println("\nBackend rodando em: http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
