package space

import (
	"context"

	"github.com/ranxx/altgotech-demo/app/request"
	"github.com/ranxx/altgotech-demo/pkg/errors"
	"github.com/ranxx/altgotech-demo/service/space"
	"github.com/ranxx/altgotech-demo/service/space/model"
)

// Space 板块
type Space struct {
}

// NewSpace new
func NewSpace(ctx context.Context) *Space {
	return &Space{}
}

// GetService new svc
func (s *Space) GetService() *space.Service {
	return space.NewService()
}

// Create 创建板块
func (s *Space) Create(ctx context.Context, uid int, req *request.SpaceCreateRequest) (*request.SpaceCreateResponse, error) {
	svc := s.GetService()
	ok, err := svc.Exist(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.NewErrCode(errors.ErrSpaceNameExist, "板块已存在")
	}

	item, err := svc.Create(ctx, uid, req.Name)
	if err != nil {
		return nil, err
	}

	return &request.SpaceCreateResponse{Item: &request.Space{
		ID:        item.ID,
		Name:      item.Name,
		CreatorID: item.CreatorID,
		AdminID:   item.AdminID,
	}}, nil
}

// ConvertFromModel ...
func ConvertFromModel(item *model.Space) *request.Space {
	return &request.Space{
		ID:        item.ID,
		Name:      item.Name,
		CreatorID: item.CreatorID,
		AdminID:   item.AdminID,
	}
}

// List 列表板块
func (s *Space) List(ctx context.Context, uid int, req *request.SpaceListRequest) (*request.SpaceListResponse, error) {
	svc := s.GetService()

	if req.All {
		uid = 0
	}

	items, err := svc.List(ctx, uid)
	if err != nil {
		return nil, err
	}

	resp := &request.SpaceListResponse{
		Items: make([]*request.Space, 0, len(items)),
	}

	for _, v := range items {
		resp.Items = append(resp.Items, ConvertFromModel(v))
	}

	return resp, nil
}

// Join 加入板块
func (s *Space) Join(ctx context.Context, uid int, req *request.SpaceJoinRequest) (*request.SpaceJoinResponse, error) {
	svc := s.GetService()
	ok, err := svc.ExistRelation(ctx, uid, int(req.ID))
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.NewErrCode(errors.ErrSpaceJoined, "已加入该板块, 不可重复加入")
	}

	space, _ := svc.Get(ctx, int(req.ID))
	if space != nil && space.CreatorID == uid {
		return nil, errors.NewErrCode(errors.ErrSpaceNoJoinSelf, "不可加入自己创建的板块")
	}

	if err = svc.CreateRelation(ctx, uid, int(req.ID)); err != nil {
		return nil, err
	}

	return &request.SpaceJoinResponse{}, nil
}

// Admin 设置管理员
func (s *Space) Admin(ctx context.Context, uid int, req *request.SpaceAdminRequest) (*request.SpaceAdminResponse, error) {
	svc := s.GetService()
	item, err := svc.Get(ctx, int(req.ID))
	if err != nil {
		return nil, err
	}
	// 如果不是创建人 ，则没有此权限
	if item.CreatorID != uid {
		return nil, errors.NewErrCode(errors.ErrSpaceNoOwner, "您不是该板块的拥有者, 无权设置管理员")
	}

	if err = svc.UpdateAdmin(ctx, int(req.ID), int(req.UID)); err != nil {
		return nil, err
	}

	return &request.SpaceAdminResponse{}, nil
}
