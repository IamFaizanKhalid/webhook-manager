package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

func router() *chi.Mux {
	// Initialize router
	r := chi.NewRouter()

	// Use middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(middleware.Logger)

	// Define CORS options
	r.Use(cors.Handler(cors.Options{
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin", "Cache-Control"},
		ExposedHeaders:   []string{"Content-Type", "JWT-Token", "Content-Disposition"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Define endpoints
	r.Route("/hooks", func(r chi.Router) {
		r.Get("/", getHooks)
		r.Post("/", createHook)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", getHook)
			r.Put("/", updateHook)
			r.Delete("/", deleteHook)
		})
	})

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/*", fs)

	return r
}
