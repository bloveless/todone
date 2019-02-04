package webapp

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type Server struct {
	Storage Storage
}

type WelcomeMessage struct{
	Welcome string `json:"welcome"`
}

// HealthCheck responds with a 200 and a body when requested.
func (s Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Healthy"))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

// Index ...
func (s Server) HomeIndex(w http.ResponseWriter, r *http.Request) {
	js, mErr := json.Marshal(WelcomeMessage{Welcome: "Home"})
	if mErr != nil {
		http.Error(w, mErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) GetAllToDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	toDos, sErr := s.Storage.GetAllToDos()
	if sErr != nil {
		http.Error(w, sErr.Error(), http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(toDos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) ToggleToDoComplete(w http.ResponseWriter, r *http.Request) {
	toDoId := chi.URLParam(r, "id")
	toggleErr := s.Storage.ToggleToDoComplete(toDoId)
	if toggleErr != nil {
		http.Error(w, toggleErr.Error(), http.StatusInternalServerError)
		return
	}

	toDos, sErr := s.Storage.GetAllToDos()
	if sErr != nil {
		http.Error(w, sErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(toDos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) ToggleToDoActive(w http.ResponseWriter, r *http.Request) {
	toDoId := chi.URLParam(r, "id")
	toggleErr := s.Storage.ToggleToDoActive(toDoId)
	if toggleErr != nil {
		http.Error(w, toggleErr.Error(), http.StatusInternalServerError)
		return
	}

	toDos, sErr := s.Storage.GetAllToDos()
	if sErr != nil {
		http.Error(w, sErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(toDos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
