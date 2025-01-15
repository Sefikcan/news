package routes

import (
	"github.com/gorilla/mux"
	"github.com/sefikcan/news/internal/handlers"
	"net/http"
)

func RegisterRoutes(router *mux.Router, newsHandler *handlers.NewsHandler) {
	router.HandleFunc("/api/news", newsHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/news", newsHandler.GetNews).Methods(http.MethodGet)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			return
		}
	}).Methods(http.MethodGet)
}
