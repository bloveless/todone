package page

import (
	"github.com/go-chi/chi"
	"net/http"
)

func AddRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/health-check", healthCheck)
	r.Get("/", home)

	return r
}
