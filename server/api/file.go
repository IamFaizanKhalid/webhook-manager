package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Static struct {
	fileServer http.Handler
}

func NewStatic(path string) *Static {
	return &Static{fileServer: http.FileServer(http.Dir(path))}
}

func (api *Static) Routes(r chi.Router) {
	r.Handle("/*", api.fileServer)
}

func (api *Static) AuthRequired() bool {
	return false
}
