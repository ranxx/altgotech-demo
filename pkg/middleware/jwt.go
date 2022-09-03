package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ranxx/altgotech-demo/config"
	"github.com/ranxx/altgotech-demo/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func jwtGinAuth(ctx *gin.Context) error {
	token := ctx.Request.Header.Get("Authorization")
	if token == "null" || token == "" {
		return errors.NewErrCode(errors.ErrAuthVerify, "请求未携带token，无权限访问")
	}
	data, err := ParseToken(token)
	if err != nil {
		return errors.NewErrCode(errors.ErrAuthVerify, "token验证失败")
	}
	idstr := data["id"]
	if idstr == nil {
		return errors.NewErrCode(errors.ErrAuthVerify, "token验证失败")
	}
	id, err := strconv.ParseInt(idstr.(string), 0, 64)
	if err != nil {
		return errors.NewErrCode(errors.ErrAuthVerify, "token验证失败")
	}
	ctx.Set("user.id", id)
	return err
}

// JwtGinAuth jwt
func JwtGinAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := jwtGinAuth(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, errors.Convert2Response(err))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// GetUserIDFromCtx 从ctx拿取 user_id
func GetUserIDFromCtx(ctx context.Context) (int64, error) {
	userid := ctx.Value("user.id")
	if userid == nil {
		return 0, errors.NewCode(errors.ErrAuthVerify)
	}
	id, ok := userid.(int64)
	if !ok {
		return 0, errors.NewCode(errors.ErrAuthVerify)
	}
	return id, nil
}

// GenerateToken token
func GenerateToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  strconv.FormatInt(id, 10),
		"exp": time.Now().Add(time.Hour * 27 * 7).Unix(), // 可以添加过期时间
	})
	tokenString, err := token.SignedString([]byte(config.GetConf().Jwttoken))
	// Bearer
	return fmt.Sprintf("Bearer %s", tokenString), err
}

// ParseToken token
func ParseToken(tokenString string) (map[string]interface{}, error) {
	// Bearer
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetConf().Jwttoken), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return claims, err
}
