package page

import (
	"encoding/json"
	"net/http"
)

type WelcomeMessage struct{
	Welcome string `json:"welcome"`
}

func home(w http.ResponseWriter, r *http.Request) {
	js, mErr := json.Marshal(WelcomeMessage{Welcome: "Home"})
	if mErr != nil {
		http.Error(w, mErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Model", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Healthy"))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}
