package routers

import (
	"commonlib/gin"
)

//权限确认
func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		//TODO 可依据登录用户cookie判断是否有权限
		context.Next()
	}
}
