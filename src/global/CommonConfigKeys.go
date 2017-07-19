//APPCONFIG.JSON配置文件KEY 说明及定义
package global

const (
	DB_CONN string = "MySqlCon" //数据库配置项名称
	IS_DEV  string = "IsDev"    //是否开发环境配置项名称

	IpWhiteList string = "IpWhiteList" //IP 白名单控制

	ENCRYPT_COOKIE_SECRET string = "EncryptCookieSecret" //COOKIE AES加密串
	TASK_URL_SECRET       string = "TaskUrlSecret"       //生成URL加密参数串

	//REDIS连接串
	REDIS_HOST string = "RedisHost"
	REDIS_DB   string = "RedisDB"
	REDIS_PWD  string = "RedisPwd"

	//日志路径
	ErrorLogPath     string = "ErrorLogPath"     //错误日志路径
	InfoLogPath      string = "InfoLogPath"      //普通信息日志路径
	WarnLogPath      string = "WarnLogPath"      //警告提醒日志
	UserAgentLogPath string = "UserAgentLogPath" //打开问卷头部信息日志

	//邮箱相关
	SmtpServerSocket    string = "SmtpServerSocket"    //服务器发送端口
	SmtpServer          string = "SmtpServer"          //邮箱服务器地址
	MailSender          string = "MailSender"          //发件人
	SenderPwd           string = "SenderPwd"           //密码
	SystemErrorReminder string = "SystemErrorReminder" //系统知会人邮件配置

	//短信发送相关接口
	MSG_App_ID  string = "MSG_App_ID"
	MSG_App_Key string = "MSG_App_Key"

	//图形接口插件字体配置
	FONT_HEITI_PATH string = "FONT_HEITI_PATH" //字体路径
)
