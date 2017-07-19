package main

import (
	"SampleProj/routers"
	Middlewares "SampleProj/routers/middlewares"
	"commonlib/gin"
	"global"
	"io/ioutil"
	"path/filepath"
	//RedisProvider "redisprovider"
	"utiltools"
)

//初始化配置文件
func InitGlobalConfigs(configFilePath string) error {
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = utiltools.FromByteJson(content, &global.ConfigMappings)
	if err != nil {
		return err
	}
	return nil
}

//初始化配置文件=>绝对路径
func init() {
	//获取配置文件路径
	appConfigJsonPath, _ := filepath.Abs("./appconfig.json")
	InitGlobalConfigs(appConfigJsonPath) //初始化配置文件映射表
	//RedisProvider.InitRedisPool()        //初始化REDIS连接池
	utiltools.StartSendAdminEmails() //启动管理员邮件队列线程
}

//API入口
func main() {
	//初始化日志logger
	global.ErrorLogger, global.InfoLogger, global.WarnLogger, global.UserAgentLogger = utiltools.InitLoggers()
	defer global.ErrorLogger.Close()
	defer global.InfoLogger.Close()
	defer global.WarnLogger.Close()
	defer global.UserAgentLogger.Close()

	if global.ConfigMappings[global.IS_DEV] == "false" {
		gin.SetMode(gin.ReleaseMode)
	}
	//初始化路由信息
	r := gin.New()
	r.Use(Middlewares.AccessLogger())       //访问日志记录
	r.Use(Middlewares.AccessRecovery())     //全局请求异常处理
	r.Use(Middlewares.IpWhiteListChecker()) //ip白名单限制
	routers.RouterBindings(r)               //自定义路由绑定
	r.Run("0.0.0.0:2016")                   //打开监听端口
}
