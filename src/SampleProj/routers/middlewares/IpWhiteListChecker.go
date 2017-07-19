package routers

import (
	entity "SampleProj/entities"
	"commonlib/gin"
	"global"
	"net/http"
	"strings"
	"utiltools"
)

//IP访问控制
func IpWhiteListChecker() gin.HandlerFunc {
	return func(context *gin.Context) {
		ip := context.ClientIP()                            //获取客户请求地址，默认是只能从NGINX服务器代理转发过来，用户不能直接在本地发起AJAX请求
		ipList := global.ConfigMappings[global.IpWhiteList] //获取配置文件中IP列表
		if !strings.Contains(ipList, ip) {
			//写日志
			utiltools.LogWarn("IP请求无效：" + ip + ",详细:" + utiltools.ToJson(context.Request.Header) + "," +
				context.Request.UserAgent() + "," + context.Request.URL.RequestURI() + ",API:" + context.Request.URL.Path)
			resultModel := entity.BadRequestResultModel("请求被忽略")
			context.String(http.StatusOK, resultModel.ToJson())
			context.Abort()
		} else {
			context.Next()
		}
	}
}
