package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
	return nil, fmt.Errorf("登录")
}

// Register 注册
func (u *User) Register(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("注册")
}
