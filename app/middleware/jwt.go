package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"venturan/app/common/response"
	"venturan/app/services"
	"venturan/global"
	"venturan/global/serviceErrors"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.Fail(c, serviceErrors.TokenIsExpired)
			c.Abort()
			return
		}
		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil {
			response.Fail(c, serviceErrors.TokenIsExpired)
			c.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.Fail(c, serviceErrors.TokenIsExpired)
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
