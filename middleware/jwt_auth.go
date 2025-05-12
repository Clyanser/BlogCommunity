package middleware

import (
	"GoBlog/models/ctype"
	"GoBlog/models/res"
	"GoBlog/service/redis_ser"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()

			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}
		//判断是否在redis中
		if redis_ser.CheckLogout(token) {
			res.FailWithMsg("token已经失效", c)
			c.Abort()
			return
		}
		//登录的用户_判断
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMsg("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}
		if redis_ser.CheckLogout(token) {
			res.FailWithMsg("token已经失效", c)
			c.Abort()
			return
		}
		//	登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMsg("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
