package todo

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (h *Handler) AddRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", h.getAllToDos)
	r.Post("/{id}/toggle-complete", h.toggleToDoComplete)
	r.Post("/{id}/toggle-active", h.toggleToDoActive)

	return r
}
