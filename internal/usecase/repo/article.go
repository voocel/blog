package repo

import (
	"blog/internal/entity"
	"context"
	"gorm.io/gorm"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

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

func (a ArticleRepo) GetArticlesRepo(ctx context.Context, page, pageSize int) ([]*entity.Article, int64, error) {
	var total int64
	var articles []*entity.Article
	err := a.db.WithContext(ctx).Scopes(Paginate(page, pageSize)).Find(&articles).Count(&total).Error
	return articles, total, err
}

func (a ArticleRepo) UpdateArticleRepo(ctx context.Context, article *entity.Article) error {
	return a.db.Updates(article).Error
}

func (a ArticleRepo) DeleteArticleRepo(ctx context.Context, aid int64) error {
	return a.db.Delete(&entity.Article{}, aid).Error
}

func (a ArticleRepo) DeleteArticleListRepo(ctx context.Context, aids []int64) error {
	return a.db.Where("id in (?)", aids).Delete(&entity.Article{}).Error
}
