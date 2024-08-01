package usecase

import (
	"blog/internal/entity"
	"context"
	"github.com/jinzhu/copier"
	"golang.org/x/sync/singleflight"
)

type ArticleUseCase struct {
	repo ArticleRepo
	sf   singleflight.Group
}

func NewArticleUseCase(r ArticleRepo) *ArticleUseCase {
	return &ArticleUseCase{repo: r}
}

func (c *ArticleUseCase) CreateArticle(ctx context.Context, req entity.ArticleReq) error {
	article := new(entity.Article)
	copier.Copy(article, req)
	_, err := c.repo.AddArticleRepo(ctx, article)
	return err
}
