package elasticsearch

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Url string
}

func NewElasticsearchClient(url string) *Client {
	return &Client{Url: url}
}

func (es *Client) Search(index, query string) ([]byte, error) {
	resp, err := http.Post(es.Url+index+"/_search", "application/json", bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Elasticsearch error: " + resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
