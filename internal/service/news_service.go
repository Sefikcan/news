package service

import (
	"github.com/sefikcan/news/internal/models"
	"github.com/sefikcan/news/internal/repository"
)

type NewsService struct {
	Repo *repository.NewsRepository
}

func NewNewsService(repository *repository.NewsRepository) *NewsService {
	return &NewsService{
		Repo: repository,
	}
}

func (s *NewsService) Create(news *models.News) error {
	return s.Repo.Create(news)
}

func (s *NewsService) Update(news *models.News) error {
	return s.Repo.Update(news)
}

func (s *NewsService) Delete(id string) error {
	return s.Repo.Delete(id)
}

func (s *NewsService) GetNews(query string) ([]models.News, error) {
	return s.Repo.GetNews(query)
}
