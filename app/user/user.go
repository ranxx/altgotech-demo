package user

import (
	"context"
	"crypto/md5"

	"github.com/ranxx/altgotech-demo/app/request"
	"github.com/ranxx/altgotech-demo/pkg/errors"
	"github.com/ranxx/altgotech-demo/pkg/middleware"
	"github.com/ranxx/altgotech-demo/service/user"
)

// User 用户
type User struct {
}

// NewUser new
func NewUser(ctx context.Context) *User {
	return &User{}
}

// passwd 密码处理
func (s *User) passwd(p string) string {
	hash := md5.New()
	hash.Write([]byte(p))
	return string(hash.Sum(nil))
}

func (u *User) GetService() *user.Service {
	return user.NewService()
}

// Login 登录
func (u *User) Login(ctx context.Context, req *request.UserLoginRequest) (*request.UserLoginResponse, error) {
	// 获取
	svc := u.GetService()
	user, err := svc.Get(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	// 判断密码对不
	if user.Passwd != u.passwd(req.Passwd) {
		return nil, errors.NewErrCode(errors.ErrUserPasswd, "账号或者密码不正确")
	}

	// 生成token
	token, err := middleware.GenerateToken(int64(user.ID))
	if err != nil {
		return nil, err
	}

	return &request.UserLoginResponse{Token: token}, nil
}

// Register 注册
func (u *User) Register(ctx context.Context, req *request.UserRegisterRequest) (*request.UserRegisterResponse, error) {
	// 检查是否注册
	svc := u.GetService()
	ok, err := svc.Exist(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.NewErrCode(errors.ErrUserEmailExist, "用户已存在")
	}

	// 注册
	if err := svc.Create(ctx, req.Email, req.Passwd); err != nil {
		return nil, err
	}

	return &request.UserRegisterResponse{}, nil
}
