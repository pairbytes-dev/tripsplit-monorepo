package main

import (
	"log"

	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/core/user"
	"github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/db"
	httpapi "github.com/pairbytes-dev/tripsplit-monorepo/backend/internal/http"
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

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
