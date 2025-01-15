package handlers

import (
	"encoding/json"
	"github.com/sefikcan/news/internal/models"
	"github.com/sefikcan/news/internal/service"
	"net/http"
	"time"
)

type NewsHandler struct {
	service *service.NewsService
}

func NewNewsHandler(service *service.NewsService) *NewsHandler {
	return &NewsHandler{
		service: service,
	}
}

func (h *NewsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var news models.News
	news.CreatedAt = time.Now()
	err := json.NewDecoder(r.Body).Decode(&news)
	if err != nil {
		return
	}
	if err := h.service.Create(&news); err != nil {
		http.Error(w, "News could not be created.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *NewsHandler) GetNews(w http.ResponseWriter, _ *http.Request) {
	query := `{
  "query": {
    "match_all": {}
  },
  "size": 10,
  "from": 0
}`
	newsList, err := h.service.GetNews(query)
	if err != nil {
		http.Error(w, "News not found!", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(newsList)
	if err != nil {
		return
	}
}
