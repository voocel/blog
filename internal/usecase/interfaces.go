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
		AddArticleRepo(ctx context.Context, article *entity.Article) error
		GetArticleByIdRepo(ctx context.Context, aid int64) (*entity.Article, error)
		GetArticlesRepo(ctx context.Context) ([]*entity.Article, error)
		DeleteArticleRepo(ctx context.Context, aid int64) error
		DeleteArticleListRepo(ctx context.Context, aids []int64) error
	}

	CategoryRepo interface {
		AddCategoryRepo(ctx context.Context, category *entity.Category) error
		GetCategoryByIdRepo(ctx context.Context, cid int64) (*entity.Category, error)
		GetCategoryByNameRepo(ctx context.Context, name string) (*entity.Category, error)
		GetCategoryByNameExistRepo(ctx context.Context, name string) (bool, error)
		GetCategoriesRepo(ctx context.Context) ([]*entity.Category, error)
	}

	StarRepo interface {
		AddStarRepo(ctx context.Context, uid, articleId int64) error
		DeleteStarRepo(ctx context.Context, uid, articleId int64) error
		GetStarsByUidRepo(ctx context.Context, uid int64) ([]*entity.Star, error)
		GetStarsByArticleIdRepo(ctx context.Context, articleId int64) ([]*entity.Star, error)
	}

	AdvertRepo interface {
		AddAdvertRepo(ctx context.Context, advert *entity.Advert) error
		DetailRepo(ctx context.Context, id int64) (*entity.Advert, error)
		GetAdvertListRepo(ctx context.Context) ([]*entity.Advert, error)
		UpdateAdvertRepo(ctx context.Context, advert *entity.Advert) error
		DeleteAdvertRepo(ctx context.Context, id int64) error
		DeleteAdvertBatchRepo(ctx context.Context, ids []int64) error
	}

	MenuRepo interface {
		AddMenuRepo(ctx context.Context, menu *entity.Menu) error
		GetMenuByIdRepo(ctx context.Context, id int64) (*entity.Menu, error)
		GetMenusRepo(ctx context.Context) ([]*entity.Menu, error)
		UpdateMenuRepo(ctx context.Context, menu *entity.Menu) error
		DeleteMenuRepo(ctx context.Context, id int64) error
		DeleteMenusBatchRepo(ctx context.Context, ids []int64) error
	}

	BannerRepo interface {
		AddBannerRepo(ctx context.Context, banner *entity.Banner) error
		GetBannerByIdRepo(ctx context.Context, id int64) (*entity.Banner, error)
		GetBannersRepo(ctx context.Context) ([]*entity.Banner, error)
		UpdateBannerRepo(ctx context.Context, banner *entity.Banner) error
		DeleteBannerRepo(ctx context.Context, id int64) error
		DeleteBannersBatchRepo(ctx context.Context, ids []int64) error
	}
)
