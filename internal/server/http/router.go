package http

import (
	"github.com/IamFaizanKhalid/webhook-api/internal/auth"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (srv *Server) buildRouter(auth *auth.Auth, addMiddlewares ...func(http.Handler) http.Handler) *chi.Mux {
	r := chi.NewRouter()

	// Config
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(30 * time.Second))
	for _, mdlwr := range addMiddlewares {
		if mdlwr != nil {
			r.Use(mdlwr)
		}
	}

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
		if api.AuthRequired() && auth != nil {
			r.With(auth.HttpMiddleware).Group(api.Routes)
		} else {
			r.Group(api.Routes)
		}
	}

	return r
}
