package space

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
	return nil, fmt.Errorf("创建板块")
}

// Join 加入板块
func (s *Space) Join(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("加入板块")
}

// Admin 设置管理员
func (s *Space) Admin(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("设置管理员")
}
