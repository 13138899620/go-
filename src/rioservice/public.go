package rioservice

import (
	"fmt"
	"global"
)

//发送邮件服务
//记录日志
func SendMail(sender, receiver, title, content string) string {
	//构造邮件内容并校验
	var mail []byte
	var err error
	if mail, err = buildMail(sender, receiver, title, content); err != nil {
		return err.Error()
	}
	//发送邮件
	err = postData(global.ConfigMappings[global.RIO_EMAIL_URL], mail)
	logContent := fmt.Sprintf("sender:%s,receiver:%s,title:%s,content:%s;error:%+v", sender, receiver, title, content, err)
	return logContent
}

//发送myoa待办
func SendMyoa(guid, formUrl, handler, startTime, dueTime, periodName, objName, roleName string) string {
	//构造邮件内容并校验
	var data []byte
	var err error
	if data, err = buildMyoaCreateBody(guid, formUrl, handler, startTime, dueTime, periodName, objName, roleName); err != nil {
		return err.Error()
	}
	//发送邮件
	err = postMyoaData(global.ConfigMappings[global.MYOA_DOMAIN]+global.ConfigMappings[global.MYOA_APPID]+"/myoa/workitem/create", data)
	logContent := fmt.Sprintf("guid:%s,formUrl:%s,handler:%s,startTime:%s;dueTime:%s;periodName:%s;objName:%s;roleName:%s;error:%+v", guid, formUrl, handler, startTime, dueTime, periodName, objName, roleName, err)
	return logContent
}

//撤回myoa待办
func DelMyoa(guid, handler string) string {
	//构造邮件内容并校验
	var data []byte
	var err error
	if data, err = buildMyoaDelBody(guid, handler); err != nil {
		return err.Error()
	}
	//发送邮件
	err = postMyoaData(global.ConfigMappings[global.MYOA_DOMAIN]+global.ConfigMappings[global.MYOA_APPID]+"/myoa/workitem/close", data)
	logContent := fmt.Sprintf("guid:%s,handler:%s,error:%+v", guid, handler, err)
	return logContent
}
