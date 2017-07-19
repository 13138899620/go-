package controllers

import (
	UserEntities "SampleProj/entities/user"
	"commonlib/gin"
	"global"
	"time"
	"utiltools"
)

//从COOKIE中获取用户信息
func GetLoginUserFromCookie(context *gin.Context) *UserEntities.User {
	//返回相应的登录信息
	cookie, _ := context.Request.Cookie(global.COOKIE_GO_USER)
	if cookie == nil {
		return nil
	}
	key := global.ConfigMappings[global.ENCRYPT_COOKIE_SECRET]
	val := utiltools.AesDecrypter(cookie.Value, []byte(key)) //解密当前登录用户
	loginUser := &UserEntities.User{}
	err := utiltools.FromJson(val, loginUser)
	if err != nil {
		return nil
	}
	return loginUser
}

//获取用户票据信息
func GetTicketByUserId(userId int64) *UserEntities.UserTicket {
	//TODO 可以将用户票据信息存储至REDIS
	//每个用户的票据信息获取都不一样，并且有过期时间限制
	//这里为了演示，所有用户都先返回一个固定的TICKET
	//新增用户签名信息
	ticket := &UserEntities.UserTicket{}
	ticket.UId = userId
	ticket.SecretKey = "vnfjur39efnvjbi2"                //16位密钥
	ticket.SigToken = "a620e03bf1924259b3b700004bf45a24" //32位签名TOKEN
	ticket.ExpiredTime = time.Now().AddDate(0, 0, 1)     //1天后过期
	ticket.CreateTime = time.Now()
	return ticket
}

//获取POST ITEM参数,统一进行POST数据解密与反序列化
func GetPostParamObj(context *gin.Context, ticket *UserEntities.UserTicket, outObj interface{}) {
	item := context.PostForm(global.POST_PARAM_NAME)
	if item != "" {
		//解密传输数据
		decryptJson := utiltools.AesDecrypter(item, []byte(ticket.SecretKey))
		utiltools.FromJson(decryptJson, &outObj)
	}
}
