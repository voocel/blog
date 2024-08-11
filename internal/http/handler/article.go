package handler

import (
	"blog/internal/entity"
	"blog/internal/repository/redis"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ArticleHandler struct {
	logUsecase     *usecase.LogstashUseCase
	articleUsecase *usecase.ArticleUseCase
	bannerUsecase  *usecase.BannerUseCase
	redis          *redis.Redis
}

type IdListReq struct {
	IDList []int64 `json:"id_list"`
}

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

func NewArticleHandler(u *usecase.ArticleUseCase, b *usecase.BannerUseCase, log *usecase.LogstashUseCase) *ArticleHandler {
	return &ArticleHandler{
		logUsecase:     log,
		articleUsecase: u,
		bannerUsecase:  b,
		redis:          redis.GetClient(),
	}
}

func (h *ArticleHandler) Create(c *gin.Context) {
	req := entity.ArticleReq{}
	resp := new(ApiResponse)
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}
	if err := h.articleUsecase.CreateArticle(c, req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	articles, total, err := h.articleUsecase.GetList(c, page, pageSize)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}
	resp.Data = articles
	resp.Paging = &PagingInformation{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) Detail(c *gin.Context) {
	resp := new(ApiResponse)
	aid := c.Param("aid")
	articleId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}

	article, err := h.articleUsecase.GetDetailById(c, int64(articleId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}
	resp.Data = article
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) Update(c *gin.Context) {
	req := entity.ArticleUpdateReq{}
	resp := new(ApiResponse)
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}

	if req.BannerID > 0 {
		banner, err := h.bannerUsecase.Detail(c, req.BannerID)
		if err != nil {
			resp.Code = 1
			resp.Message = err.Error()
			c.JSON(http.StatusOK, resp)
			h.logUsecase.AddLogstash(c, &entity.Logstash{
				Level:     entity.ErrorLevel,
				Content:   err.Error(),
				CreatedAt: time.Now(),
			})
			return
		}
		req.BannerUrl = banner.Path
	}
	if err := h.articleUsecase.UpdateArticle(c, req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) DeleteArticleById(c *gin.Context) {
	resp := new(ApiResponse)
	aid := c.Param("aid")
	articleId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}

	if err := h.articleUsecase.DeleteArticle(c, int64(articleId)); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) DeleteArticlesBatch(c *gin.Context) {
	resp := new(ApiResponse)
	var req IdListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}

	if err := h.articleUsecase.DeleteArticles(c, req.IDList); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) Like(c *gin.Context) {
	resp := new(ApiResponse)
	aid := c.Param("aid")
	articleId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
		return
	}
	if err := h.redis.Hincrby(c, "article_like", strconv.Itoa(articleId), 1); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		h.logUsecase.AddLogstash(c, &entity.Logstash{
			Level:     entity.ErrorLevel,
			Content:   err.Error(),
			CreatedAt: time.Now(),
		})
	}
	resp.Message = "点赞成功!"
	c.JSON(http.StatusOK, resp)
	return
}

func (h *ArticleHandler) Calendar(c *gin.Context) {
	resp := new(ApiResponse)
	var data = make([]CalendarResponse, 0)
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0)
	days := int(now.Sub(aYearAgo).Hours() / 24)
	for i := 0; i < days; i++ {
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")
		data = append(data, CalendarResponse{
			Date:  day,
			Count: 1,
		})
	}
	resp.Data = data
	c.JSON(http.StatusOK, resp)
}
