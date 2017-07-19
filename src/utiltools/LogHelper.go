/*日志记录工具*/
package utiltools

import (
	log "commonlib/log4go"
	"fmt"
	"global"
)

//管理员邮件内容
type adminEmail struct {
	Title   string
	Content string
}

var adminEmailChan chan *adminEmail //管理员提醒邮件内存缓冲队列

//初始化邮件队列
func init() {
	adminEmailChan = make(chan *adminEmail, 512)
}

//初始化日志LOGGERS
func InitLoggers() (*log.Logger, *log.Logger, *log.Logger, *log.Logger) {
	//初始化信息ERROR LOGGER
	errWriter := log.NewFileWriter()
	errWriter.SetPathPattern(global.ConfigMappings[global.ErrorLogPath])
	errorLogger := log.NewLogger()
	errorLogger.Register(errWriter)
	errorLogger.SetLevel(log.ERROR)
	//初始化INFO LOGER
	infoWriter := log.NewFileWriter()
	infoWriter.SetPathPattern(global.ConfigMappings[global.InfoLogPath])
	infoLogger := log.NewLogger()
	infoLogger.Register(infoWriter)
	infoLogger.SetLevel(log.INFO)
	//初始化警告日志
	warnWriter := log.NewFileWriter()
	warnWriter.SetPathPattern(global.ConfigMappings[global.WarnLogPath])
	warnLogger := log.NewLogger()
	warnLogger.Register(warnWriter)
	warnLogger.SetLevel(log.WARNING)
	//初妈化USERAGENT日志
	userAgentWriter := log.NewFileWriter()
	userAgentWriter.SetPathPattern(global.ConfigMappings[global.UserAgentLogPath])
	userAgentlogger := log.NewLogger()
	userAgentlogger.Register(userAgentWriter)
	userAgentlogger.SetLevel(log.INFO)
	return errorLogger, infoLogger, warnLogger, userAgentlogger
}

//发送管理员邮件
func StartSendAdminEmails() {
	go func() {
		for {
			item, ok := <-adminEmailChan
			if !ok {
				LogError("管理员邮件提醒队列异常")
			} else {
				SendToAdmin(item.Title, item.Content)
			}
		}
	}()
}

//警告日志
func LogWarn(content string) {
	go func(content string) {
		if global.ConfigMappings[global.IS_DEV] == "true" {
			fmt.Println(content)
		}
		global.WarnLogger.Warn("%s", content)
	}(content)
}

//记录错误日志
func LogError(content string) {
	go func(content string) {
		if global.ConfigMappings[global.IS_DEV] == "true" {
			fmt.Println(content)
		}
		global.ErrorLogger.Error("%s", content)
		email := &adminEmail{}
		email.Title = "YUN ASSESS云评估系统提醒"
		email.Content = content
		adminEmailChan <- email //先推送到队列
	}(content)
}

//记录信息访问日志
func LogInfo(content string) {
	go func(content string) {
		if global.ConfigMappings[global.IS_DEV] == "true" {
			fmt.Println(content)
		}
		global.InfoLogger.Info("%s", content)
	}(content)
}

//记录打开问卷客户端信息
func LogUserAgentInfo(content string) {
	go func(content string) {
		if global.ConfigMappings[global.IS_DEV] == "true" {
			fmt.Println(content)
		}
		global.UserAgentLogger.Info("%s", content)
	}(content)
}
