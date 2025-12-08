package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"fitranker-api/internal/http/handler"
)

func New() http.Handler {
	r := chi.NewRouter()

	// ヘルスチェック
	r.Get("/health", handler.Health)

	// 今後ここに /users や /auth などを追加していく
	// r.Route("/users", func(r chi.Router) { ... })

	return r
}
