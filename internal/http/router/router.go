package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"fitranker-api/internal/http/handler"
	"fitranker-api/internal/training"
)

func New(trainingService training.Service) http.Handler {
	r := chi.NewRouter()

	// ヘルスチェック
	r.Get("/health", handler.Health)

	tc := handler.NewTrainingController(trainingService)

	r.Route("/api", func(r chi.Router) {
		r.Get("/personal/{id}", tc.GetPersonalInfo)
		r.Get("/ranking", tc.GetRanking)
		r.Post("/training-records", tc.PostTrainingRecords)
	})

	return r
}
