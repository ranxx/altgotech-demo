package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranxx/altgotech-demo/pkg/components"
	"github.com/ranxx/altgotech-demo/pkg/errors"
)

// EncodeGinResponse ...
type EncodeGinResponse func(ctx *gin.Context) (interface{}, error)

// Gin wrap
func (e EncodeGinResponse) Gin(ctx *gin.Context) {
	resp, err := e(ctx)
	if err != nil {
		components.ServerErrorEncoder(ctx, err, ctx.Writer)
		ctx.Abort()
		return
	}
	components.EncodeHTTPGenericResponse(ctx, ctx.Writer, resp)
}

// EncodeGinMiddle ...
type EncodeGinMiddle func(ctx *gin.Context) error

// Gin wrap
func (e EncodeGinMiddle) Gin(ctx *gin.Context) {
	err := e(ctx)
	if err == nil {
		return
	}
	components.ServerErrorEncoder(ctx, err, ctx.Writer)
	ctx.Abort()
}

// GinRecover panic
func GinRecover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				e, ok := err.(error)
				if ok {
					components.ServerErrorEncoder(ctx, e, ctx.Writer)
				} else {
					components.ServerErrorEncoder(ctx, errors.NewErrCode(errors.Err, fmt.Sprintf("%v", err)), ctx.Writer)
				}
				ctx.Abort()
				return
			}
		}()
		ctx.Next()
	}
}

// GinCors 跨域设置
func GinCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		ctx.Next()
	}
}
