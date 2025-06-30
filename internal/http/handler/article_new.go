package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandlerNew struct {
	articleUsecase *usecase.ArticleUseCase
	tagUsecase     *usecase.TagUseCase
}

func NewArticleHandlerNew(articleUsecase *usecase.ArticleUseCase, tagUsecase *usecase.TagUseCase) *ArticleHandlerNew {
	return &ArticleHandlerNew{
		articleUsecase: articleUsecase,
		tagUsecase:     tagUsecase,
	}
}

// GetArticles 获取文章列表
func (h *ArticleHandlerNew) GetArticles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// category := c.Query("category")
	// tag := c.Query("tag")
	// search := c.Query("search")
	// status := c.Query("status")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	articles, total, err := h.articleUsecase.GetList(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, err.Error()))
		return
	}

	var articleResponses []entity.ArticleResponse
	for _, article := range articles {
		articleResponses = append(articleResponses, convertToArticleResponse(article))
	}

	paginatedData := entity.NewPaginatedResponse(articleResponses, int(total), page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// GetArticle 获取文章详情
func (h *ArticleHandlerNew) GetArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "文章ID格式错误"))
		return
	}

	article, err := h.articleUsecase.GetDetailById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, entity.NewErrorResponse(404, "文章不存在"))
		return
	}

	response := convertToArticleResponse(article)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(response, "获取成功"))
}

// CreateArticle 创建文章
func (h *ArticleHandlerNew) CreateArticle(c *gin.Context) {
	var req entity.ArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	err := h.articleUsecase.CreateArticle(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// UpdateArticle 更新文章
func (h *ArticleHandlerNew) UpdateArticle(c *gin.Context) {
	idStr := c.Param("id")
	_, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "文章ID格式错误"))
		return
	}

	var req entity.ArticleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	err = h.articleUsecase.UpdateArticle(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "更新成功"))
}

// DeleteArticle 删除文章
func (h *ArticleHandlerNew) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "文章ID格式错误"))
		return
	}

	err = h.articleUsecase.DeleteArticle(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

func convertToArticleResponse(article *entity.Article) entity.ArticleResponse {
	return entity.ArticleResponse{
		ID:           strconv.FormatInt(article.ID, 10),
		Title:        article.Title,
		Subtitle:     article.Subtitle,
		Content:      article.Content,
		Excerpt:      article.Excerpt,
		CoverImage:   article.CoverImage,
		Status:       article.Status,
		IsOriginal:   article.IsOriginal,
		ViewCount:    article.ViewCount,
		LikeCount:    article.LikeCount,
		CommentCount: article.CommentCount,
		CreatedAt:    article.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    article.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
