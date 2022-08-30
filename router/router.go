package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ranxx/altgotech-demo/pkg/middleware"
	"github.com/ranxx/altgotech-demo/router/user"
)

func Router() http.Handler {
	engine := gin.Default()

	api := engine.Group("/api/v1")

	userCtl := user.NewUser()
	api.GET("/user/login", middleware.EncodeGinResponse(userCtl.Login).Gin)
	api.GET("/user/register", middleware.EncodeGinResponse(userCtl.Register).Gin)

	// blogCtl := blog.NewBlog()

	return engine
}
