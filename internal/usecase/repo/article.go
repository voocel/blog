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

	err := a.db.WithContext(ctx).Model(&entity.Article{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = a.db.WithContext(ctx).Scopes(Paginate(page, pageSize)).Find(&articles).Error
	return articles, total, err
}

func (a ArticleRepo) UpdateArticleRepo(ctx context.Context, article *entity.Article) error {
	return a.db.WithContext(ctx).Where("id = ?", article.ID).Updates(article).Error
}

func (a ArticleRepo) DeleteArticleRepo(ctx context.Context, aid int64) error {
	return a.db.Delete(&entity.Article{}, aid).Error
}

func (a ArticleRepo) DeleteArticleListRepo(ctx context.Context, aids []int64) error {
	return a.db.Where("id in (?)", aids).Delete(&entity.Article{}).Error
}

func (a ArticleRepo) IncrementViewCountRepo(ctx context.Context, aid int64) error {
	return a.db.WithContext(ctx).Model(&entity.Article{}).Where("id = ?", aid).
		UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}

func (a ArticleRepo) AddArticleTagsRepo(ctx context.Context, articleId int64, tagIds []int64) error {
	var articleTags []entity.ArticleTag
	for _, tagId := range tagIds {
		articleTags = append(articleTags, entity.ArticleTag{
			ArticleID: articleId,
			TagID:     tagId,
		})
	}
	return a.db.WithContext(ctx).Create(&articleTags).Error
}

func (a ArticleRepo) DeleteArticleTagsRepo(ctx context.Context, articleId int64) error {
	return a.db.WithContext(ctx).Where("article_id = ?", articleId).Delete(&entity.ArticleTag{}).Error
}

func (a ArticleRepo) GetArticleTagsRepo(ctx context.Context, articleId int64) ([]int64, error) {
	var articleTags []entity.ArticleTag
	err := a.db.WithContext(ctx).Where("article_id = ?", articleId).Find(&articleTags).Error
	if err != nil {
		return nil, err
	}

	var tagIds []int64
	for _, at := range articleTags {
		tagIds = append(tagIds, at.TagID)
	}
	return tagIds, nil
}
