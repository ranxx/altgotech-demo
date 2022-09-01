package blog

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Blog 博客
type Blog struct {
}

// NewBlog ...
func NewBlog() *Blog {
	return &Blog{}
}

// FullTextSearch 全文搜索
func (b *Blog) FullTextSearch(ctx *gin.Context) (interface{}, error) {
	return nil, fmt.Errorf("全文搜索")
}
