package main

import (
	"log"

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

	conn, err := db.OpenPostgres(cfg)
	if err != nil {
		log.Fatalf("could not connect to postgres: %v", err)
	}

	defer conn.Close()

	router := httpapi.NewRouter(conn)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
