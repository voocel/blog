package handler

import (
	"blog/internal/entity"
	"blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	CommentUseCase *usecase.CommentUseCase
}

func NewCommentHandler(u *usecase.CommentUseCase) *CommentHandler {
	return &CommentHandler{
		CommentUseCase: u,
	}
}

func (h *CommentHandler) Create(c *gin.Context) {
	resp := new(ApiResponse)
	var req entity.CommentReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}
	err = h.CommentUseCase.AddComment(c, &entity.Comment{
		ArticleID:       req.ArticleID,
		Content:         req.Content,
		ParentCommentID: req.ParentCommentID,
		UserID:          0,
	})
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
	return
}

func (h *CommentHandler) GetArticleCommentList(c *gin.Context) {
	resp := new(ApiResponse)
	aid := c.Param("aid")
	articleId, err := strconv.Atoi(aid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	comments, err := h.CommentUseCase.GetCommentsByArticleId(c, int64(articleId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = comments
	c.JSON(http.StatusOK, resp)
	return
}

func (h *CommentHandler) GetAllCommentList(c *gin.Context) {
	resp := new(ApiResponse)
	comments, err := h.CommentUseCase.GetComments(c)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = comments
	c.JSON(http.StatusOK, resp)
	return
}

// Delete 删除评论及子评论
func (h *CommentHandler) Delete(c *gin.Context) {
	resp := new(ApiResponse)
	cid := c.Param("cid")
	commentId, err := strconv.Atoi(cid)
	if err != nil {
		resp.Code = 1
		resp.Message = "params invalid"
		c.JSON(http.StatusOK, resp)
		return
	}

	// 查询评论
	comment, err := h.CommentUseCase.GetCommentById(c, int64(commentId))
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	user, exists := c.Get("jwt-user")
	u, ok := user.(*entity.User)
	if !exists || !ok {
		resp.Code = 1
		resp.Message = "token invalid"
		c.JSON(http.StatusOK, resp)
		return
	}
	// 必须本人或管理员才可删除
	if u.ID != comment.UserID || u.Role != 1 {
		resp.Code = 1
		resp.Message = "permission denied"
		c.JSON(http.StatusOK, resp)
		return
	}
	err = h.CommentUseCase.DeleteComment(c, int64(commentId))
	if err != nil {
		resp.Code = 1
	}
}
