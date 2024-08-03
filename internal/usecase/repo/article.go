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

func (a ArticleRepo) AddArticleRepo(ctx context.Context, article *entity.Article) error {
	return a.db.Create(article).Error
}

func (a ArticleRepo) GetArticleByIdRepo(ctx context.Context, aid int64) (*entity.Article, error) {
	var model = new(entity.Article)
	err := a.db.First(model, aid).Error
	return model, err
}

func (a ArticleRepo) GetArticlesRepo(ctx context.Context) ([]*entity.Article, error) {
	var articles []*entity.Article
	err := a.db.Find(&articles).Error
	return articles, err
}

func (a ArticleRepo) DeleteArticleRepo(ctx context.Context, aid int64) error {
	return a.db.Delete(&entity.Article{}, aid).Error
}

func (a ArticleRepo) DeleteArticleListRepo(ctx context.Context, aids []int64) error {
	return a.db.Where("id in (?)", aids).Delete(&entity.Article{}).Error
}
