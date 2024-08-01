package usecase

import (
	"blog/internal/entity"
	"context"
)

type (
	UserRepo interface {
		GetUserByIdRepo(ctx context.Context, uid int64) (*entity.User, error)
		GetUserByNameRepo(ctx context.Context, name string) (*entity.User, error)
		GetUserByNameExistRepo(ctx context.Context, name string) (bool, error)
		GetUsersRepo(ctx context.Context) ([]*entity.User, int, error)
		AddUserRepo(ctx context.Context, user *entity.User) (*entity.User, error)
		UpdateAvatarUserRepo(ctx context.Context, uid int64, avatar string) (int, error)
		UpdateAddressUserRepo(ctx context.Context, uid int64, address string) (int, error)
	}

	ArticleRepo interface {
		AddArticleRepo(ctx context.Context, article *entity.Article) (*entity.Article, error)
		GetArticleByIdRepo(ctx context.Context, aid int64) (*entity.Article, error)
		GetArticlesRepo(ctx context.Context) ([]*entity.Article, int, error)
	}
)
