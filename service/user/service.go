package user

import (
	"context"

	"github.com/ranxx/altgotech-demo/pkg/db"
	"github.com/ranxx/altgotech-demo/service/user/dao"
	"github.com/ranxx/altgotech-demo/service/user/model"
)

// Service user
type Service struct {
	d *dao.User
}

// NewService new
func NewService() *Service {
	return &Service{
		d: dao.NewUser(db.Mysql()),
	}
}

// Exist 是否存在
func (s *Service) Exist(ctx context.Context, email string) (bool, error) {
	return s.d.Exist(ctx, email)
}

// Create 创建用户
func (s *Service) Create(ctx context.Context, email, passwd string) error {
	return s.d.Create(ctx, &model.User{
		Email:  email,
		Passwd: passwd,
	})
}

// Get 获取用户
func (s *Service) Get(ctx context.Context, email string) (*model.User, error) {
	return s.d.Get(ctx, email)
}
