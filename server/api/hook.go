package api

import (
	"github.com/IamFaizanKhalid/webhook-api/internal/server/http/api/request"
	"github.com/IamFaizanKhalid/webhook-api/internal/server/http/api/response"
	"github.com/IamFaizanKhalid/webhook-api/server/dao"
	"github.com/IamFaizanKhalid/webhook-api/server/logic"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Hook struct {
	svc *logic.CoreLogic
}

func NewHook(svc *logic.CoreLogic) *Hook {
	return &Hook{svc: svc}
}

func (api *Hook) Routes(r chi.Router) {
	r.Route("/hooks", func(r chi.Router) {
		r.Get("/", api.GetHooks)
		r.Post("/", api.CreateHook)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", api.GetHook)
			r.Put("/", api.UpdateHook)
			r.Delete("/", api.DeleteHook)
		})
	})
}

func (api *Hook) GetHooks(w http.ResponseWriter, r *http.Request) {
	response.Encode(w, api.svc.GetAllHooks(r.Context()))
}

func (api *Hook) GetHook(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetParam(r, "id")
	if err != nil {
		response.EncodeErr(w, err)
		return
	}

	h, err := api.svc.GetHook(r.Context(), id)
	if err != nil {
		response.EncodeErr(w, err)
		return
	}

	response.Encode(w, h)
}

func (api *Hook) CreateHook(w http.ResponseWriter, r *http.Request) {
	var h dao.Hook
	err := request.Decode(r, &h)
	if err != nil {
		response.EncodeErr(w, err)
		return
	}

	err = api.svc.AddHook(r.Context(), h)
	if err != nil {
		log.Println(err)
		response.EncodeErr(w, err)
		return
	}

	response.Encode(w, "hook added")
}

func (api *Hook) UpdateHook(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetParam(r, "id")
	if err != nil {
		response.EncodeErr(w, err)
		return
	}

	var h dao.Hook
	err = request.Decode(r, &h)
	if err != nil {
		response.EncodeErr(w, err)
		return
	}

	err = api.svc.UpdateHook(r.Context(), id, h)
	if err != nil {
		log.Println(err)
		response.EncodeErr(w, err)
		return
	}

	response.Encode(w, "hook updated")
}

func (api *Hook) DeleteHook(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetParam(r, "id")
	if err != nil {
		response.EncodeErr(w, err)
		return
	}

	err = api.svc.DeleteHook(r.Context(), id)
	if err != nil {
		log.Println(err)
		response.EncodeErr(w, err)
		return
	}

	response.Encode(w, "hook deleted")
}
