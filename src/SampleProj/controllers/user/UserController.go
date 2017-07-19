package controllers

import (
	BaseController "SampleProj/controllers"
	Entities "SampleProj/entities"
	UserEntities "SampleProj/entities/user"
	Facade "SampleProj/facade"
	"commonlib/gin"

	"net/http"
	"strings"
	"utiltools"
)

//HTTP METHOD:GET
//获取用户签名和加解密KEY
func GetUserTicket(context *gin.Context) {
	//1. 从COOKIE中获取用户ID
	loginUser := BaseController.GetLoginUserFromCookie(context)
	if loginUser == nil {
		context.String(http.StatusInternalServerError, "用户身份获取失败")
	} else {
		ticket := BaseController.GetTicketByUserId(loginUser.Id) //获取用户票据信息
		result := Entities.OkResultModel(utiltools.ToJson(ticket))
		context.String(http.StatusOK, utiltools.ToJson(result)) //返回用户TOKEN
	}
}

//HTTP METHOD:GET
//获取所有用户列表
func GetUserList(context *gin.Context) {
	result := Facade.GetUserList()
	context.String(http.StatusOK, result.ToJson())
}

//HTTP METHOD:POST
//参数：entities.User{} json字符串
//添加用户
func AddUser(context *gin.Context) {
	jsonStr := context.PostForm("item") //获取POST数据
	var user UserEntities.User
	err := utiltools.FromJson(jsonStr, &user)
	if err != nil {
		result := Entities.BadRequestResultModel("提交数据错误")
		context.String(http.StatusOK, result.ToJson())
	} else {
		if strings.TrimSpace(user.UserName) == "" {
			result := Entities.BadRequestResultModel("提交数据错误")
			context.String(http.StatusOK, result.ToJson())
		} else {
			result := Facade.AddUser(&user)
			context.String(http.StatusOK, result.ToJson())
		}
	}
}

//加密获取所有用户列表
//HTTP METHOD:GET
//获取所有用户列表
func GetEncryptedUserList(context *gin.Context) {
	loginUser := BaseController.GetLoginUserFromCookie(context)
	if loginUser == nil {
		context.String(http.StatusInternalServerError, "用户身份获取失败")
	} else {
		result := Facade.GetUserList()
		ticket := BaseController.GetTicketByUserId(loginUser.Id)
		context.String(http.StatusOK, result.ToEncryptJson([]byte(ticket.SecretKey)))
	}
}

//HTTP METHOD:POST
//参数：entities.User{} json字符串加密数据
//接口添加数字签名验证，见routers middlewares CheckSignature
//添加用户
func AddEncryptedUser(context *gin.Context) {
	loginUser := BaseController.GetLoginUserFromCookie(context)
	if loginUser == nil {
		context.String(http.StatusInternalServerError, "用户身份获取失败")
	} else {
		userTicket := BaseController.GetTicketByUserId(loginUser.Id)
		var user UserEntities.User
		//统一获取相应的参数
		BaseController.GetPostParamObj(context, userTicket, &user)
		//判断数据的完整性，再保存
		if strings.TrimSpace(user.UserName) == "" {
			result := Entities.BadRequestResultModel("提交数据错误")
			context.String(http.StatusOK, result.ToJson())
		} else {
			result := Facade.AddUser(&user)
			context.String(http.StatusOK, result.ToJson())
		}
	}
}
