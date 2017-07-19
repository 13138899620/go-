//路由定义
//具体可以参考GIN　开源库配置方式
//URL: https://github.com/gin-gonic/gin
package routers

import (
	AccountController "SampleProj/controllers/account"
	UserController "SampleProj/controllers/user"
	Middlewares "SampleProj/routers/middlewares"
	"commonlib/gin"
	"net/http"
)

//对请求进行分组管理，各组可以采用不同的中间件进行数据校验，权限验证
func RouterBindings(router *gin.Engine) {
	//避免返回请求404
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.String(http.StatusOK, "")
	})

	//用户操作登录模块
	account := router.Group("/account")
	{
		account.GET("/Login", AccountController.Login)
		account.GET("/LogOut", AccountController.LogOut)
		//TODO 可以直接在指定API请求再添加相应的中间件过滤请求
		//account.POST("/LogOnCompanyAdmin", Middlewares.CheckIsAdmin(), Account.LogOnCompanyAdmin)
	}

	//用户基础信息模块
	user := router.Group("/user")
	{
		user.GET("/GetUserTicket", UserController.GetUserTicket)               //获取用户票据信息
		user.GET("/GetUserList", UserController.GetUserList)                   //获取用户列表
		user.GET("/GetEncryptedUserList", UserController.GetEncryptedUserList) //传输加密用户列表

		user.POST("/AddUser", UserController.AddUser)                                                 //添加用户信息
		user.POST("/AddEncryptedUser", Middlewares.CheckSignature(), UserController.AddEncryptedUser) //添加用户信息－数据加密
	}

}
