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

func (c *ArticleUseCase) CreateArticle(ctx context.Context, req entity.ArticleRequest) error {
	article := new(entity.Article)
	copier.Copy(article, req)
	err := c.repo.AddArticleRepo(ctx, article)
	return err
}

func (c *ArticleUseCase) GetDetailById(ctx context.Context, aid int64) (*entity.Article, error) {
	return c.repo.GetArticleByIdRepo(ctx, aid)
}

func (c *ArticleUseCase) GetList(ctx context.Context, page, pageSize int) ([]*entity.Article, int64, error) {
	return c.repo.GetArticlesRepo(ctx, page, pageSize)
}

func (c *ArticleUseCase) UpdateArticle(ctx context.Context, req entity.ArticleUpdateRequest) error {
	article := new(entity.Article)
	copier.Copy(article, req)
	err := c.repo.UpdateArticleRepo(ctx, article)
	return err
}

func (c *ArticleUseCase) DeleteArticle(ctx context.Context, aid int64) error {
	return c.repo.DeleteArticleRepo(ctx, aid)
}

func (c *ArticleUseCase) DeleteArticles(ctx context.Context, aids []int64) error {
	return c.repo.DeleteArticleListRepo(ctx, aids)
}
