//全局变量
package global

import (
	log "commonlib/log4go"
	"commonlib/redigo/redis"
)

//配置文件映射表
var ConfigMappings map[string]string

//Redis连接池
var RedisClient *redis.Pool

//日志记录
var ErrorLogger *log.Logger     //错误日志LOGGER
var InfoLogger *log.Logger      //消息日志LOGGER
var WarnLogger *log.Logger      //警告提醒日志Logger
var UserAgentLogger *log.Logger //用户请求头－UserAgent日志Logger
