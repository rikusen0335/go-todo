package api

import (
	"log/slog"
	"net/http"

	"github.com/rikusen0335/go-todo/internal/pkg/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
    logger.Init()
    
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

    slog.Info("Server has started")

    http.ListenAndServe(":8000", r)
}