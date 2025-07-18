package usecase

import (
	"blog/internal/entity"
	"context"
)

type (
	UserRepo interface {
		GetUserByIdRepo(ctx context.Context, uid int64) (*entity.User, error)
		GetUserByNameRepo(ctx context.Context, name string) (*entity.User, error)
		GetUserByEmailRepo(ctx context.Context, email string) (*entity.User, error)
		GetUserByNameExistRepo(ctx context.Context, name string) (bool, error)
		GetUserByEmailExistRepo(ctx context.Context, email string) (bool, error)
		GetUsersRepo(ctx context.Context, page, pageSize int, search string) ([]*entity.User, int64, error)
		AddUserRepo(ctx context.Context, user *entity.User) error
		UpdateUserRepo(ctx context.Context, user *entity.User) error
		DeleteUserRepo(ctx context.Context, uid int64) error
	}

	ArticleRepo interface {
		AddArticleRepo(ctx context.Context, article *entity.Article) error
		UpdateArticleRepo(ctx context.Context, article *entity.Article) error
		GetArticleByIdRepo(ctx context.Context, aid int64) (*entity.Article, error)
		GetArticlesRepo(ctx context.Context, page, pageSize int) ([]*entity.Article, int64, error)
		DeleteArticleRepo(ctx context.Context, aid int64) error
		DeleteArticleListRepo(ctx context.Context, aids []int64) error
		IncrementViewCountRepo(ctx context.Context, aid int64) error

		// ArticleTag相关方法
		AddArticleTagsRepo(ctx context.Context, articleId int64, tagIds []int64) error
		DeleteArticleTagsRepo(ctx context.Context, articleId int64) error
		GetArticleTagsRepo(ctx context.Context, articleId int64) ([]int64, error)
	}

	CategoryRepo interface {
		AddCategoryRepo(ctx context.Context, category *entity.Category) error
		GetCategoryByIdRepo(ctx context.Context, cid int64) (*entity.Category, error)
		GetCategoryByNameRepo(ctx context.Context, name string) (*entity.Category, error)
		GetCategoryByNameExistRepo(ctx context.Context, name string) (bool, error)
		GetCategoriesRepo(ctx context.Context) ([]*entity.Category, error)
		UpdateCategoryRepo(ctx context.Context, category *entity.Category) error
		DeleteCategoryRepo(ctx context.Context, cid int64) error
		GetCategoryByPathExistRepo(ctx context.Context, path string) (bool, error)
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
		AddMenuRepo(ctx context.Context, menu *entity.Menu) (*entity.Menu, error)
		GetMenuByIdRepo(ctx context.Context, id int64) (*entity.Menu, error)
		GetMenuByPathRepo(ctx context.Context, path string) (*entity.Menu, error)
		GetMenusRepo(ctx context.Context) ([]*entity.Menu, error)
		UpdateMenuRepo(ctx context.Context, menu *entity.Menu) error
		DeleteMenuRepo(ctx context.Context, id int64) error
		DeleteMenusBatchRepo(ctx context.Context, ids []int64) error
		IsTitlePathExistRepo(ctx context.Context, title, path string) bool
	}

	BannerRepo interface {
		AddBannerRepo(ctx context.Context, banner *entity.Banner) error
		GetBannerByIdRepo(ctx context.Context, id int64) (*entity.Banner, error)
		GetBannersRepo(ctx context.Context) ([]*entity.Banner, error)
		UpdateBannerRepo(ctx context.Context, banner *entity.Banner) error
		DeleteBannerRepo(ctx context.Context, id int64) error
		DeleteBannersBatchRepo(ctx context.Context, ids []int64) error
	}

	CommentRepo interface {
		AddCommentRepo(ctx context.Context, comment *entity.Comment) error
		GetCommentByIdRepo(ctx context.Context, id int64) (*entity.Comment, error)
		GetCommentsByArticleIdRepo(ctx context.Context, aid int64) ([]*entity.Comment, error)
		GetCommentsRepo(ctx context.Context) ([]*entity.Comment, error)
		GetCommentsWithPaginationRepo(ctx context.Context, page, pageSize int, articleId, discussionId *int64) ([]*entity.Comment, int64, error)
		UpdateCommentRepo(ctx context.Context, comment *entity.Comment) error
		DeleteCommentRepo(ctx context.Context, id int64) error
	}

	TagRepo interface {
		AddTagRepo(ctx context.Context, tag *entity.Tag) error
		AddTagsRepo(ctx context.Context, tags []*entity.Tag) error
		GetTagByIdRepo(ctx context.Context, id int64) (*entity.Tag, error)
		GetTagByNameRepo(ctx context.Context, name string) (*entity.Tag, error)
		GetTagByNameExistRepo(ctx context.Context, name string) (bool, error)
		GetTagsRepo(ctx context.Context) ([]*entity.Tag, error)
		UpdateTagRepo(ctx context.Context, tag *entity.Tag) error
		DeleteTagRepo(ctx context.Context, id int64) error
		DeleteTagsBatchRepo(ctx context.Context, ids []int64) error
	}

	LogstashRepo interface {
		AddLogstashRepo(ctx context.Context, logstash *entity.Logstash) error
		GetLogstashByIdRepo(ctx context.Context, id int64) (*entity.Logstash, error)
		GetLogstashRepo(ctx context.Context) ([]*entity.Logstash, error)
		DeleteLogstashRepo(ctx context.Context, id int64) error
		DeleteLogstashBatchRepo(ctx context.Context, ids []int64) error
	}

	MenuBannerRepo interface {
		AddMenuBannerRepo(ctx context.Context, menuBanner *entity.MenuBanner) error
		AddMenuBannerBatchRepo(ctx context.Context, menuBanners []*entity.MenuBanner) error
		GetMenuBannerByIdRepo(ctx context.Context, id int64) (*entity.MenuBanner, error)
		GetMenuBannerByMenuIdRepo(ctx context.Context, menuId int64) ([]*entity.MenuBanner, error)
		GetMenuBannerByBannerIdRepo(ctx context.Context, bannerId int64) ([]*entity.MenuBanner, error)
		GetMenuBannersRepo(ctx context.Context) ([]*entity.MenuBanner, error)
		UpdateMenuBannerRepo(ctx context.Context, menuBanner *entity.MenuBanner) error
		DeleteMenuBannerRepo(ctx context.Context, id int64) error
	}

	DiscussionRepo interface {
		AddDiscussionRepo(ctx context.Context, discussion *entity.Discussion) error
		GetDiscussionByIdRepo(ctx context.Context, id int64) (*entity.Discussion, error)
		GetDiscussionsRepo(ctx context.Context, page, pageSize int, tagId *int64, search string) ([]*entity.Discussion, int64, error)
		UpdateDiscussionRepo(ctx context.Context, discussion *entity.Discussion) error
		DeleteDiscussionRepo(ctx context.Context, id int64) error
		IncrementViewCountRepo(ctx context.Context, id int64) error

		// 讨论标签关联
		AddDiscussionTagsRepo(ctx context.Context, discussionId int64, tagIds []int64) error
		DeleteDiscussionTagsRepo(ctx context.Context, discussionId int64) error
		GetDiscussionTagsRepo(ctx context.Context, discussionId int64) ([]int64, error)
		GetDiscussionsByTagIdRepo(ctx context.Context, tagId int64) ([]*entity.Discussion, error)
	}

	ReplyRepo interface {
		AddReplyRepo(ctx context.Context, reply *entity.Reply) error
		GetReplyByIdRepo(ctx context.Context, id int64) (*entity.Reply, error)
		GetRepliesByDiscussionIdRepo(ctx context.Context, discussionId int64) ([]*entity.Reply, error)
		UpdateReplyRepo(ctx context.Context, reply *entity.Reply) error
		DeleteReplyRepo(ctx context.Context, id int64) error
		GetReplyCountByDiscussionIdRepo(ctx context.Context, discussionId int64) (int64, error)
	}

	FriendlinkRepo interface {
		AddFriendlinkRepo(ctx context.Context, friendlink *entity.FriendLink) error
		GetFriendlinkByIdRepo(ctx context.Context, id int64) (*entity.FriendLink, error)
		GetFriendlinksRepo(ctx context.Context, page, pageSize int, status string) ([]*entity.FriendLink, int64, error)
		UpdateFriendlinkRepo(ctx context.Context, friendlink *entity.FriendLink) error
		DeleteFriendlinkRepo(ctx context.Context, id int64) error
	}

	StatisticsRepo interface {
		GetUsersCountRepo(ctx context.Context) (int, error)
		GetArticlesCountRepo(ctx context.Context) (int, error)
		GetCommentsCountRepo(ctx context.Context) (int, error)
		GetVisitsCountRepo(ctx context.Context) (int, error)
	}
)
