package es

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func New() *elasticsearch.TypedClient {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		panic(err)
	}
	return client
}
