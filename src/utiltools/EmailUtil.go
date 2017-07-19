/*外网邮件发送*/
package utiltools

import (
	"commonlib/email"
	"global"
	"net/smtp"
	"strings"
)

//发送邮件给系统管理员
func SendToAdmin(title string, content string) {
	admins := global.ConfigMappings[global.SystemErrorReminder]
	receivers := strings.Split(admins, ";")
	bcc := []string{}
	cc := []string{}
	SendMail(title, content, receivers, bcc, cc)
}

//发送邮SendMail
//title:标题
//content:邮件正文
//receivers:收件人
//bcc:密送
//cc:抄送
func SendMail(title string, content string, receivers []string, bcc []string, cc []string) bool {
	socket := global.ConfigMappings[global.SmtpServerSocket] //server:port
	sender := global.ConfigMappings[global.MailSender]       //发件人
	pwd := global.ConfigMappings[global.SenderPwd]           //客户端密码
	server := global.ConfigMappings[global.SmtpServer]       //SMTP服务器
	//创建邮件
	e := email.NewEmail()
	e.From = "YunAssess <" + sender + ">"
	e.To = receivers
	//密送
	if len(bcc) > 0 {
		e.Bcc = bcc
	}
	//抄送
	if len(cc) > 0 {
		e.Cc = cc
	}
	e.Subject = title        //邮件标题
	e.HTML = []byte(content) //采用HTML格式进行邮件发送
	err := e.Send(socket, smtp.PlainAuth("", sender, pwd, server))
	if err != nil {
		//发送失败，要记录一下日志
		LogWarn("发送邮件失败：" + ToJson(e) + ";错误：" + err.Error())
		return false
	}
	return true
}
