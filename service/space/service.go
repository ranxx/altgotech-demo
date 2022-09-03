package space

import (
	"context"

	"github.com/ranxx/altgotech-demo/pkg/db"
	"github.com/ranxx/altgotech-demo/service/space/dao"
	"github.com/ranxx/altgotech-demo/service/space/model"
)

// Service space
type Service struct {
	d *dao.Space
}

// NewService new
func NewService() *Service {
	return &Service{
		d: dao.NewSpace(db.Mysql()),
	}
}

// Exist 是否存在
func (s *Service) Exist(ctx context.Context, name string) (bool, error) {
	return s.d.Exist(ctx, name)
}

// ExistRelation 是否加入过 板块
func (s *Service) ExistRelation(ctx context.Context, uid, id int) (bool, error) {
	return s.d.ExistRelation(ctx, uid, id)
}

// Create 创建板块
func (s *Service) Create(ctx context.Context, uid int, name string) (*model.Space, error) {
	item := &model.Space{
		Name:      name,
		CreatorID: uid,
	}
	return item, s.d.Create(ctx, item)
}

// List 列表板块
func (s *Service) List(ctx context.Context, uid int) ([]*model.Space, error) {
	return s.d.List(ctx, uid)
}

// All 列表板块
func (s *Service) All(ctx context.Context) ([]*model.Space, error) {
	return s.d.List(ctx, -1)
}

// CreateRelation 创建
func (s *Service) CreateRelation(ctx context.Context, uid, id int) error {
	return s.d.CreateRelation(ctx, &model.SpaceUser{
		SID:    id,
		UserID: uid,
	})
}

// UpdateAdmin 更新管理员
func (s *Service) UpdateAdmin(ctx context.Context, id, admin int) error {
	return s.d.UpdateAdmin(ctx, id, admin)
}

// Get 获取板块
func (s *Service) Get(ctx context.Context, id int) (*model.Space, error) {
	return s.d.Get(ctx, id)
}
