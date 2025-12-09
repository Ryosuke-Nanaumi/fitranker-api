package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"fitranker-api/internal/http/router"
	"fitranker-api/internal/training"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal("open database failed")
	}
	defer func() {
		err := db.Close()
		if err != nil {

		}
	}()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Database connected!")
	repository := training.NewRepository(db)
	service := training.NewService(repository)
	r := router.New(service)

	log.Printf("🚀 Server running at http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
