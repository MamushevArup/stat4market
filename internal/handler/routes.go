package handler

import (
	_ "github.com/MamushevArup/stat4market/docs"
	"github.com/MamushevArup/stat4market/internal/handler/event"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/ping"))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:4444/swagger/doc.json"),
	))

	r.Post("/api/event", event.Save(h.repository.Clickhouse))
	return r
}
