package usecase

import (
	"blog/internal/entity"
	"blog/pkg/log"
	"context"

	"github.com/jinzhu/copier"
	"golang.org/x/sync/singleflight"
)

type ArticleUseCase struct {
	repo         ArticleRepo
	userRepo     UserRepo
	categoryRepo CategoryRepo
	tagRepo      TagRepo
	sf           singleflight.Group
}

func NewArticleUseCase(r ArticleRepo, userRepo UserRepo, categoryRepo CategoryRepo, tagRepo TagRepo) *ArticleUseCase {
	return &ArticleUseCase{
		repo:         r,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
		tagRepo:      tagRepo,
	}
}

func (c *ArticleUseCase) CreateArticle(ctx context.Context, req entity.ArticleRequest, userID int64) error {
	article := new(entity.Article)
	copier.Copy(article, req)

	// 设置用户ID
	article.UserID = userID

	// 创建文章
	err := c.repo.AddArticleRepo(ctx, article)
	if err != nil {
		return err
	}

	// 处理标签关联
	if len(req.TagIds) > 0 {
		err = c.repo.AddArticleTagsRepo(ctx, article.ID, req.TagIds)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ArticleUseCase) GetDetailById(ctx context.Context, aid int64) (*entity.Article, error) {
	return c.repo.GetArticleByIdRepo(ctx, aid)
}

// GetDetailByIdWithRelations 获取文章详情（包含关联信息）
func (c *ArticleUseCase) GetDetailByIdWithRelations(ctx context.Context, aid int64) (*entity.ArticleWithRelations, error) {
	// 获取文章基本信息
	article, err := c.repo.GetArticleByIdRepo(ctx, aid)
	if err != nil {
		return nil, err
	}

	// 增加访问次数
	err = c.repo.IncrementViewCountRepo(ctx, aid)
	if err != nil {
		log.Error("Error incrementing view count:", err.Error())
	}

	// 获取更新后的文章信息
	article, err = c.repo.GetArticleByIdRepo(ctx, aid)
	if err != nil {
		return nil, err
	}

	result := &entity.ArticleWithRelations{
		Article: article,
	}

	// 获取作者信息
	if article.UserID > 0 {
		user, err := c.userRepo.GetUserByIdRepo(ctx, article.UserID)
		if err == nil {
			result.User = user
		}
	}

	// 获取分类信息
	if article.CategoryID > 0 {
		category, err := c.categoryRepo.GetCategoryByIdRepo(ctx, article.CategoryID)
		if err == nil {
			result.Category = category
		}
	}

	// 获取标签信息
	tagIds, err := c.repo.GetArticleTagsRepo(ctx, aid)
	if err == nil && len(tagIds) > 0 {
		var tags []*entity.Tag
		for _, tagId := range tagIds {
			tag, err := c.tagRepo.GetTagByIdRepo(ctx, tagId)
			if err == nil {
				tags = append(tags, tag)
			}
		}
		result.Tags = tags
	}

	return result, nil
}

func (c *ArticleUseCase) GetList(ctx context.Context, page, pageSize int) ([]*entity.Article, int64, error) {
	return c.repo.GetArticlesRepo(ctx, page, pageSize)
}

func (c *ArticleUseCase) UpdateArticle(ctx context.Context, articleId int64, req entity.ArticleUpdateRequest) error {
	article := new(entity.Article)
	article.ID = articleId
	copier.Copy(article, req)

	// 更新文章
	err := c.repo.UpdateArticleRepo(ctx, article)
	if err != nil {
		return err
	}

	// 更新标签关联
	if req.TagIds != nil {
		// 先删除旧的关联
		err = c.repo.DeleteArticleTagsRepo(ctx, articleId)
		if err != nil {
			return err
		}

		// 添加新的关联
		if len(req.TagIds) > 0 {
			err = c.repo.AddArticleTagsRepo(ctx, articleId, req.TagIds)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *ArticleUseCase) DeleteArticle(ctx context.Context, aid int64) error {
	// 删除文章标签关联
	err := c.repo.DeleteArticleTagsRepo(ctx, aid)
	if err != nil {
		return err
	}

	// 删除文章
	return c.repo.DeleteArticleRepo(ctx, aid)
}

func (c *ArticleUseCase) DeleteArticles(ctx context.Context, aids []int64) error {
	// 删除文章标签关联
	for _, aid := range aids {
		err := c.repo.DeleteArticleTagsRepo(ctx, aid)
		if err != nil {
			return err
		}
	}

	// 删除文章
	return c.repo.DeleteArticleListRepo(ctx, aids)
}
