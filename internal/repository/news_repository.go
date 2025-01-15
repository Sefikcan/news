package repository

import (
	"encoding/json"
	"github.com/sefikcan/news/internal/couchbase"
	"github.com/sefikcan/news/internal/elasticsearch"
	"github.com/sefikcan/news/internal/models"
)

type NewsRepository struct {
	Couchbase     *couchbase.Client
	Elasticsearch *elasticsearch.Client
}

func NewNewsRepository(couchbase *couchbase.Client, elasticsearch *elasticsearch.Client) *NewsRepository {
	return &NewsRepository{
		Couchbase:     couchbase,
		Elasticsearch: elasticsearch,
	}
}

func (repo *NewsRepository) Create(news *models.News) error {
	_, err := repo.Couchbase.Collection().Upsert(news.Id, news, nil)
	return err
}

func (repo *NewsRepository) Update(news *models.News) error {
	_, err := repo.Couchbase.Collection().Upsert(news.Id, news, nil)
	return err
}

func (repo *NewsRepository) Delete(id string) error {
	_, err := repo.Couchbase.Collection().Remove(id, nil)
	return err
}

func (repo *NewsRepository) GetNews(query string) ([]models.News, error) {
	result, err := repo.Elasticsearch.Search("news", query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Hits struct {
			Hits []struct {
				Source models.News `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.Unmarshal(result, &response); err != nil {
		return nil, err
	}

	newsList := []models.News{}
	for _, hit := range response.Hits.Hits {
		newsList = append(newsList, hit.Source)
	}

	return newsList, nil
}
