package usecase

import (
	"blog/internal/entity"
	"context"
	"github.com/gin-gonic/gin"
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
	err := c.repo.AddArticleRepo(ctx, article)
	return err
}

func (c *ArticleUseCase) GetDetailById(ctx context.Context, id int) (*entity.Article, error) {
	return c.repo.GetArticleByIdRepo(ctx, int64(id))
}

func (c *ArticleUseCase) GetList(ctx *gin.Context) ([]*entity.Article, error) {
	return c.repo.GetArticlesRepo(ctx)
}
