package article

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ranxx/altgotech-demo/app/article"
	"github.com/ranxx/altgotech-demo/app/request"
	"github.com/ranxx/altgotech-demo/pkg/errors"
	"github.com/ranxx/altgotech-demo/pkg/middleware"
)

// Article ...
type Article struct {
}

// NewArticle ...
func NewArticle() *Article {
	return &Article{}
}

// Post 发表文章
func (a *Article) Post(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 取参数
	id, err := strconv.ParseInt(ctx.Param("sid"), 10, 64)
	if err != nil {
		return nil, err
	}

	req := new(request.ArticlePostRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}

	req.Space = int(id)

	resp, err := article.NewArticle(ctx).Post(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Comment 发表评论
func (a *Article) Comment(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 取参数
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		return nil, err
	}

	req := new(request.ArticleCommentRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}
	req.AID = int(aid)

	resp, err := article.NewArticle(ctx).Comment(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Like 点赞
func (a *Article) Like(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 取参数
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		return nil, err
	}

	req := new(request.ArticleLikeRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}
	req.AID = int(aid)

	resp, err := article.NewArticle(ctx).Like(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UnLike 取消点赞
func (a *Article) UnLike(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 取参数
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		return nil, err
	}

	req := new(request.ArticleUnLikeRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}
	req.AID = int(aid)

	resp, err := article.NewArticle(ctx).UnLike(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Top 置顶
func (a *Article) Top(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 取参数
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		return nil, err
	}

	req := new(request.ArticleTopRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}
	req.AID = int(aid)

	resp, err := article.NewArticle(ctx).Top(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// List 列表
func (a *Article) List(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 取参数
	sid, err := strconv.ParseInt(ctx.Param("sid"), 10, 64)
	if err != nil {
		return nil, err
	}

	req := new(request.ArticleListRequest)
	if err := ctx.BindQuery(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}
	req.Space = int(sid)

	resp, err := article.NewArticle(ctx).List(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Detail 详情
func (a *Article) Detail(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 取参数
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		return nil, err
	}

	req := new(request.ArticleDetailRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}
	req.AID = int(aid)

	resp, err := article.NewArticle(ctx).Detail(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
