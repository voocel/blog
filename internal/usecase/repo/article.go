package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

type ArticleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	return &ArticleRepo{db: db}
}

func (a ArticleRepo) AddArticleRepo(ctx context.Context, article *entity.Article) (*entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepo) GetArticleByIdRepo(ctx context.Context, aid int64) (*entity.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a ArticleRepo) GetArticlesRepo(ctx context.Context) ([]*entity.Article, int, error) {
	//TODO implement me
	panic("implement me")
}
