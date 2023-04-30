package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (srv *Server) buildRouter() *chi.Mux {
	r := chi.NewRouter()

	// Config
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin", "Cache-Control"},
		ExposedHeaders:   []string{"Content-Type", "JWT-Token", "Content-Disposition"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Adding routes
	for _, api := range srv.apis {
		r.Group(api.Routes)
	}

	return r
}
