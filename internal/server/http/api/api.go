package api

import (
	"github.com/go-chi/chi/v5"
)

type API interface {
	Routes(r chi.Router)
	AuthRequired() bool
}
