package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	mid "github.com/ranxx/altgotech-demo/pkg/middleware"
	"github.com/ranxx/altgotech-demo/router/article"
	"github.com/ranxx/altgotech-demo/router/blog"
	"github.com/ranxx/altgotech-demo/router/space"
	"github.com/ranxx/altgotech-demo/router/user"
)

func Router() http.Handler {
	engine := gin.Default()

	api := engine.Group("/api/v1")

	userCtl := user.NewUser()
	api.POST("/user/login", mid.EncodeGinResponse(userCtl.Login).Gin)       // 登录
	api.POST("/user/register", mid.EncodeGinResponse(userCtl.Register).Gin) // 注册

	blogCtl := blog.NewBlog()
	api.POST("/blog/search", mid.EncodeGinResponse(blogCtl.FullTextSearch).Gin) // 全文搜索

	spaceCtl := space.NewSpace()
	api.POST("/space", mid.JwtGinAuth(), mid.EncodeGinResponse(spaceCtl.Create).Gin)           // 创建板块
	api.POST("/space/:sid/join", mid.JwtGinAuth(), mid.EncodeGinResponse(spaceCtl.Join).Gin)   // 加入板块
	api.POST("/space/:sid/admin", mid.JwtGinAuth(), mid.EncodeGinResponse(spaceCtl.Admin).Gin) // 设置管理员

	articleCtl := article.NewArticle()
	articleG := api.Group("/space/:sid", mid.JwtGinAuth())
	articleG.POST("/article", mid.EncodeGinResponse(articleCtl.Post).Gin)                 // 发表文章
	articleG.GET("/articles", mid.EncodeGinResponse(articleCtl.List).Gin)                 // 列表文章
	articleG.POST("/article/:aid/comment", mid.EncodeGinResponse(articleCtl.Comment).Gin) // 评论文章
	articleG.POST("/article/:aid/like", mid.EncodeGinResponse(articleCtl.Like).Gin)       // 点赞文章
	articleG.POST("/article/:aid/top", mid.EncodeGinResponse(articleCtl.Top).Gin)         // 置顶文章
	articleG.GET("/article/:aid/detail", mid.EncodeGinResponse(articleCtl.Detail).Gin)    // 详情文章

	return engine
}
