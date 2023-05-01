package api

import (
	"github.com/IamFaizanKhalid/webhook-api/internal/server/http/api/response"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Ping struct{}

func NewPing() *Ping {
	return &Ping{}
}

func (api *Ping) Routes(r chi.Router) {
	r.Get("/hello", api.Hello)
}

func (api *Ping) Hello(w http.ResponseWriter, _ *http.Request) {
	response.Encode(w, "everything is working fine..!")
}
