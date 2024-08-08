package es

import (
	"blog/internal/entity"
	"context"
	"encoding/json"
)

func index() string {
	return "article_index"
}

func detail(id string) (*entity.Article, error) {
	client := New()
	resp, err := client.Get(index(), id).Do(context.Background())
	if err != nil {
		return nil, err
	}
	model := new(entity.Article)
	err = json.Unmarshal(resp.Source_, model)
	if err != nil {
		return nil, err
	}
	//model.ID = resp.Id_
	return model, nil
}
