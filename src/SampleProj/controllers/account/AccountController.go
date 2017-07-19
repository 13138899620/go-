package account

import (
	UserEntities "SampleProj/entities/user"
	"commonlib/gin"
	"global"
	"net/http"
	"time"
	"utiltools"
)

//登录操作
//身份验证写COOKIE
func Login(context *gin.Context) {
	//TODO 判断用户身份

	//初始化默认数据
	loginUser := &UserEntities.User{}
	loginUser.Id = 37225
	loginUser.UserName = "stoneldeng"

	//全局COOKIE加密KEY
	key := global.ConfigMappings[global.ENCRYPT_COOKIE_SECRET]
	//加密登录user内容
	val := utiltools.AesEncrypter([]byte(utiltools.ToJson(loginUser)), []byte(key))
	cookie := &http.Cookie{Name: global.COOKIE_GO_USER, Value: val, Path: "/", HttpOnly: true}
	//过期时间可选，不设置默认为SESSION COOKIE类型，关闭浏览器就失效
	expires := time.Now().AddDate(0, 0, 1)
	cookie.Expires = expires //设置一天后过期
	http.SetCookie(context.Writer, cookie)
	context.String(http.StatusOK, "login success")
}

//退出操作
//清除用户登录COOKIE
func LogOut(context *gin.Context) {
	//TODO 清除COOKIE
	context.String(http.StatusOK, "log out")
}
