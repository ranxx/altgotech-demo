package article

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Article ...
type Article struct {
}

// NewArticle ...
func NewArticle() *Article {
	return &Article{}
}

// Post 发表文章
func (a *Article) Post(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("发表文章")
}

// Comment 发表评论
func (a *Article) Comment(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("发表评论")
}

// Like 点赞
func (a *Article) Like(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("点赞")
}

// Top 置顶
func (a *Article) Top(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("置顶")
}

// List 列表
func (a *Article) List(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("列表")
}

// Detail 详情
func (a *Article) Detail(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("详情")
}
