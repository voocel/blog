package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagHandler struct {
	tagUsecase *usecase.TagUseCase
}

func NewTagHandler(u *usecase.TagUseCase) *TagHandler {
	return &TagHandler{tagUsecase: u}
}

func (h *TagHandler) Create(c *gin.Context) {
	req := entity.TagReq{}
	resp := new(ApiResponse)
	if err := c.ShouldBind(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.tagUsecase.AddTag(c, req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *TagHandler) Detail(c *gin.Context) {
	resp := new(ApiResponse)
	tid := c.Param("tid")
	tagId, err := strconv.Atoi(tid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	tag, err := h.tagUsecase.GetTagById(c, int64(tagId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	resp.Data = tag
	c.JSON(http.StatusOK, resp)
	return
}

func (h *TagHandler) List(c *gin.Context) {
	resp := new(ApiResponse)
	tags, err := h.tagUsecase.GetTags(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = tags
	c.JSON(http.StatusOK, resp)
	return
}

func (h *TagHandler) Update(c *gin.Context) {
	resp := new(ApiResponse)
	var req entity.TagReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.tagUsecase.UpdateTag(c, req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *TagHandler) Delete(c *gin.Context) {
	resp := new(ApiResponse)
	tid := c.Param("tid")
	tagId, err := strconv.Atoi(tid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := h.tagUsecase.DeleteTag(c, int64(tagId)); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
}

func (h *TagHandler) DeleteBatch(c *gin.Context) {
	resp := new(ApiResponse)
	var req IdListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}
	if err := h.tagUsecase.DeleteTagBatch(c, req.IDList); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}
