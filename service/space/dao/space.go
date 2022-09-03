package dao

import (
	"context"

	"github.com/ranxx/altgotech-demo/pkg/errors"
	"github.com/ranxx/altgotech-demo/service/space/model"
	"gorm.io/gorm"
)

// Space 板块
type Space struct {
	db *gorm.DB
}

// NewSpace new space
func NewSpace(db *gorm.DB) *Space {
	return &Space{
		db: db,
	}
}

// Exist name是否存在
func (s *Space) Exist(ctx context.Context, name string) (bool, error) {
	total := (int64)(0)
	if err := s.db.Model(&model.Space{}).WithContext(ctx).Where("name = ?", name).Count(&total).Error; err != nil {
		return false, err
	}
	return total == 0, nil
}

// ExistRelation 是否加入过 板块
func (s *Space) ExistRelation(ctx context.Context, uid, id int) (bool, error) {
	total := (int64)(0)
	if err := s.db.Model(&model.SpaceUser{}).WithContext(ctx).Where("sid = ? and user_id = ?", id, uid).Count(&total).Error; err != nil {
		return false, err
	}
	return total == 0, nil
}

// Create 创建
func (s *Space) Create(ctx context.Context, item *model.Space) error {
	return s.db.WithContext(ctx).Create(item).Error
}

// List 列表板块
func (s *Space) List(ctx context.Context, uid int) ([]*model.Space, error) {
	items := make([]*model.Space, 0, 10)

	dbr := s.db.Model(&model.Space{}).WithContext(ctx)
	if uid > 0 {
		dbr = dbr.Select("space.*").Joins("LEFT JOIN space_user ON space_user.sid = space.id and space_user.user_id = ?", uid).Where("space.creator_id = ?", uid)
	}

	err := dbr.Order("space.created_at").Find(&items).Error

	return items, err
}

// CreateRelation 创建
func (s *Space) CreateRelation(ctx context.Context, item *model.SpaceUser) error {
	return s.db.WithContext(ctx).Create(item).Error
}

// UpdateAdmin 更新管理员
func (s *Space) UpdateAdmin(ctx context.Context, id, admin int) error {
	return s.db.Model(&model.Space{}).WithContext(ctx).Where("id = ?", id).Update("admin_id", admin).Error
}

// Get 获取
func (s *Space) Get(ctx context.Context, id int) (*model.Space, error) {
	item := model.Space{}

	ret := s.db.Model(&model.Space{}).WithContext(ctx).Where("id = ?", id).First(&item)
	if ret.Error != nil {
		return nil, ret.Error
	}

	if ret.RowsAffected == 0 {
		return nil, errors.NewErrCode(errors.ErrSpaceNotFound, "板块不存在")
	}

	return &item, nil
}
