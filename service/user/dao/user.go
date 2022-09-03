package dao

import (
	"context"

	"github.com/ranxx/altgotech-demo/pkg/errors"
	"github.com/ranxx/altgotech-demo/service/user/model"
	"gorm.io/gorm"
)

// User user
type User struct {
	db *gorm.DB
}

// NewUser new user
func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

// Exist email是否存在
func (u *User) Exist(ctx context.Context, email string) (bool, error) {
	total := (int64)(0)
	if err := u.db.Model(&model.User{}).WithContext(ctx).Where("email = ?", email).Count(&total).Error; err != nil {
		return false, err
	}
	return total == 0, nil
}

// Create 创建
func (u *User) Create(ctx context.Context, item *model.User) error {
	return u.db.WithContext(ctx).Create(item).Error
}

// Get 获取
func (u *User) Get(ctx context.Context, email string) (*model.User, error) {
	item := model.User{}

	ret := u.db.Model(&model.User{}).WithContext(ctx).Where("email = ?", email).First(&item)
	if ret.RowsAffected == 0 && ret.Error == gorm.ErrRecordNotFound {
		return nil, errors.NewErrCode(errors.ErrUserNotFound, "用户不存在")
	}
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &item, nil
}
