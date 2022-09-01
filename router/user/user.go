package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ranxx/altgotech-demo/app/request"
	"github.com/ranxx/altgotech-demo/app/user"
	"github.com/ranxx/altgotech-demo/pkg/errors"
)

// User 用户
type User struct {
}

// NewUser ...
func NewUser() *User {
	return &User{}
}

// Login 登录
func (u *User) Login(ctx *gin.Context) (interface{}, error) {
	req := new(request.UserLoginRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}

	resp, err := user.NewUser(ctx).Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Register 注册
func (u *User) Register(ctx *gin.Context) (interface{}, error) {
	req := new(request.UserRegisterRequest)
	if err := ctx.BindJSON(req); err != nil {
		return nil, errors.NewErrCode(errors.ErrBadRequest, err.Error())
	}

	resp, err := user.NewUser(ctx).Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
