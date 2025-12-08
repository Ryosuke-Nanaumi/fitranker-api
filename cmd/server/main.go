package main

import (
	"log"
	"net/http"
	"os"

	"fitranker-api/internal/http/router"
	"fitranker-api/internal/training"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	var trainingService training.Service
	r := router.New(trainingService)
	// TODO: service, repositoryの実装を追加

	log.Printf("🚀 Server running at http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
