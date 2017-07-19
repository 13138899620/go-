package routers

import (
	BaseController "SampleProj/controllers"
	Entities "SampleProj/entities"
	"commonlib/gin"
	"global"
	"net/http"
	"strings"
	"utiltools"
)

//验证企业管理员用户数字签名是否正确
func CheckSignature() gin.HandlerFunc {
	//初始化其它数据
	return func(context *gin.Context) {
		//获取头部信息
		uid := context.Request.Header.Get(global.UID)
		timestamp := context.Request.Header.Get(global.TIMESTAMP)
		sig := context.Request.Header.Get(global.SIGNATURE)
		postItem := context.PostForm(global.POST_PARAM_NAME)
		resultModel := &Entities.ResultModel{}
		//判断头部信息中是否包含数据
		if uid == "" || timestamp == "" || sig == "" || postItem == "" {
			utiltools.LogWarn("请求头缺失,签名验证不通过：" + utiltools.ToJson(context.Request.Header) + ",API:" + context.Request.URL.Path)
			resultModel = Entities.BadRequestResultModel("签名验证不通过，请求参数缺失")
			context.String(http.StatusOK, resultModel.ToJson())
			context.Abort()
		} else {
			loginUser := BaseController.GetLoginUserFromCookie(context)
			//TODO 1. 添加时间戳请求唯一性判断，可采用REDIS的HASH判断 uid-timestamp
			userTicket := BaseController.GetTicketByUserId(loginUser.Id)
			//TODO 2.判断userTicket是否已经期 now.after(userTicket.ExpiredTime)
			source := uid + timestamp + postItem
			//计算签名
			signature := utiltools.Sha1Sig(source, userTicket.SigToken)
			if strings.ToUpper(signature) != strings.ToUpper(sig) {
				utiltools.LogWarn("签名验证不通过：" + utiltools.ToJson(context.Request.Header) + "SERVER签名：" + signature + ",API:" + context.Request.URL.Path)
				resultModel = Entities.BadRequestResultModel("签名验证不通过")
				context.String(http.StatusOK, resultModel.ToJson())
				context.Abort()
			} else {
				//通过验证直接下一步
				context.Next()
			}

		}
	}
}
