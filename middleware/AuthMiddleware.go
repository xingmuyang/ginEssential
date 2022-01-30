package middleware

import (
	"github.com/gin-gonic/gin"
	"learn/ginEssential/common"
	"learn/ginEssential/models"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取浏览器用户token
		tokenString := ctx.GetHeader("Authorization")

		//token格式校验
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		//获取claims
		token, claims, err := common.ParseToken(tokenString)
		if err!= nil || !token.Valid{
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//从claims获取UserId
		userId := claims.UserId

		//校验数据
		DB := common.GetDB()
		var user models.User
		DB.First(&user, userId)

		// 用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()

	}
}
