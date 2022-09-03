package space

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ranxx/altgotech-demo/app/request"
	"github.com/ranxx/altgotech-demo/app/space"
	"github.com/ranxx/altgotech-demo/pkg/errors"
	"github.com/ranxx/altgotech-demo/pkg/middleware"
)

// Space 板块
type Space struct {
}

// NewSpace ...
func NewSpace() *Space {
	return &Space{}
}

// Create 创建板块
func (s *Space) Create(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req := new(request.SpaceCreateRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}

	resp, err := space.NewSpace(ctx).Create(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// List 创建板块
func (s *Space) List(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req := new(request.SpaceListRequest)
	if err := ctx.BindQuery(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}

	resp, err := space.NewSpace(ctx).List(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Join 加入板块
func (s *Space) Join(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 取参数
	id, err := strconv.ParseInt(ctx.Param("sid"), 10, 64)
	if err != nil {
		return nil, err
	}

	req := &request.SpaceJoinRequest{
		ID: id,
	}

	resp, err := space.NewSpace(ctx).Join(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Admin 设置管理员
func (s *Space) Admin(ctx *gin.Context) (interface{}, error) {
	uid, err := middleware.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// 取参数
	id, err := strconv.ParseInt(ctx.Param("sid"), 10, 64)
	if err != nil {
		return nil, err
	}

	req := new(request.SpaceAdminRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}

	req.ID = id

	resp, err := space.NewSpace(ctx).Admin(ctx, int(uid), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
