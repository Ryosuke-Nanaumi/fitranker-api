package main

import (
	"log"
	"net/http"
	"os"

	"fitranker-api/internal/http/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	r := router.New()

	log.Printf("🚀 Server running at http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
