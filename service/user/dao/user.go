package dao

import (
	"context"

	"github.com/ranxx/altgotech-demo/pkg/errors"
	"github.com/ranxx/altgotech-demo/service/user/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

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
	return u.db.Create(item).Error
}

// Get 获取
func (u *User) Get(ctx context.Context, email string) (*model.User, error) {
	item := model.User{}

	ret := u.db.Model(&model.User{}).Where("email = ?", email).First(&item)
	if ret.Error != nil {
		return nil, ret.Error
	}

	if ret.RowsAffected == 0 {
		return nil, errors.NewErrCode(errors.ErrUserNotFound, "用户不存在")
	}

	return &item, nil
}
