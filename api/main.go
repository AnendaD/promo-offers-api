package main

import (
	"context"
	"net/http"

	"promo-offers-api/internal/config"
	httpserver "promo-offers-api/internal/http"
	"promo-offers-api/internal/logger"

	"github.com/go-chi/chi"

	"github.com/go-chi/chi/middleware"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logger.New(config.AppConfig.Level)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// TODO: create handlers for every endpoint
	r.Post("/offers", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/merchants/{id}/offers", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/offers/calculate", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/offers/{id}/activate", func(w http.ResponseWriter, r *http.Request) {})

	HTTPServer := httpserver.New(config.AppConfig.Address, r, logger)
	HTTPServer.Run(ctx)
}
