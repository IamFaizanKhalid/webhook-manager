package main

import (
	"encoding/json"
	"fmt"
	"github.com/IamFaizanKhalid/webhook-api/model"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func getHooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hooks.GetAllHooks())
}

func getHook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	if id == "" {
		json.NewEncoder(w).Encode("id is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h, err := hooks.GetHook(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(h)
}

func createHook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var h model.Hook
	err := json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = hooks.AddHook(h)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("failed to add hook")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("hook added")
}

func updateHook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	if id == "" {
		json.NewEncoder(w).Encode("id is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var h model.Hook
	err := json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !hooks.Exists(id) {
		json.NewEncoder(w).Encode(fmt.Sprintf("hook with id `%s` not found", id))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = hooks.UpdateHook(id, h)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("failed to update hook")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("hook updated")
}

func deleteHook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	if id == "" {
		json.NewEncoder(w).Encode("id is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !hooks.Exists(id) {
		json.NewEncoder(w).Encode(fmt.Sprintf("hook with id `%s` not found", id))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := hooks.DeleteHook(id)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode("failed to delete hook")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("hook deleted")
}
