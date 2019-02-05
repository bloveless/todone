package todo

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type Handler struct {
	storage Storage
}

// NewHandler will return a new todo handler.
func NewHandler(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

// Index ...
func (h *Handler) HomeIndex(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) getAllToDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Model", "application/json")
	toDos, sErr := h.storage.GetAllToDos()
	if sErr != nil {
		http.Error(w, sErr.Error(), http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(toDos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) toggleToDoComplete(w http.ResponseWriter, r *http.Request) {
	toDoId := chi.URLParam(r, "id")
	toggleErr := h.storage.ToggleToDoComplete(toDoId)
	if toggleErr != nil {
		http.Error(w, toggleErr.Error(), http.StatusInternalServerError)
		return
	}

	toDos, sErr := h.storage.GetAllToDos()
	if sErr != nil {
		http.Error(w, sErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Model", "application/json")
	err := json.NewEncoder(w).Encode(toDos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) toggleToDoActive(w http.ResponseWriter, r *http.Request) {
	toDoId := chi.URLParam(r, "id")
	toggleErr := h.storage.ToggleToDoActive(toDoId)
	if toggleErr != nil {
		http.Error(w, toggleErr.Error(), http.StatusInternalServerError)
		return
	}

	toDos, sErr := h.storage.GetAllToDos()
	if sErr != nil {
		http.Error(w, sErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Model", "application/json")
	err := json.NewEncoder(w).Encode(toDos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
