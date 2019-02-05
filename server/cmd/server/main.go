package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
	"todone-backend/api/page"
	"todone-backend/api/storage"
	"todone-backend/api/todo"
)

func main() {
	inMemStorage := &storage.InMemoryStorage{}
	toDoHandler := todo.NewHandler(inMemStorage)

	r := getConfiguredRouter()
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/", page.AddRoutes())
		r.Mount("/to-do", toDoHandler.AddRoutes())
	})

	err := http.ListenAndServe(":9090", r)
	// err := http.ListenAndServeTLS(":4443", "keys/server.crt", "keys/server.key", r)
	if err != nil {
		fmt.Println(err)
	}
}

func getConfiguredRouter() *chi.Mux {
	// Creates a router without any middleware by default
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	corsMiddleware := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(corsMiddleware.Handler)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(30 * time.Second))

	return r
}
