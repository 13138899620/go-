/*
REDIS 常用操作：
1. KEY-VALUE
2. LIST
3. HASHTABLE
更多redis操作命令可参见：https://redis.io/commands
关于Redigo相关使用文档：https://github.com/garyburd/redigo
*/
package redisprovider

import (
	"commonlib/redigo/redis"
	"global"
	"strconv"
	"time"
	UtilTools "utiltools"
)

//初始化REDIS 连接池
func InitRedisPool() {
	//初始化REDIS缓存连接池配置
	REDIS_HOST := global.ConfigMappings[global.REDIS_HOST]
	REDIS_DB, _ := strconv.ParseInt(global.ConfigMappings[global.REDIS_DB], 10, 64)
	REDIS_PWD := global.ConfigMappings[global.REDIS_PWD] //REDIS 密码
	//初始化连接池
	global.RedisClient = &redis.Pool{
		MaxIdle:     20,
		MaxActive:   8000,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			//REDIS需密码访问
			if _, err = c.Do("AUTH", REDIS_PWD); err != nil {
				c.Close()
				return nil, err
			}
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

//获取指定数量的队列数据
//listName:队列名称
//count:队列元素个数
//返回:[]string
func GetRangeFromQueue(listName string, count int64) []string {
	client := global.RedisClient.Get()
	defer client.Close()
	items, err := redis.Strings(client.Do("LRANGE", listName, 0, count))
	if err != nil {
		UtilTools.LogError("获取REDIS LIST 数据失败：" + listName + ";" + err.Error())
		return nil
	}
	return items
}

//判断redis key是否存在
func ExistsKey(key string) bool {
	client := global.RedisClient.Get()
	defer client.Close()
	exists, err := redis.Bool(client.Do("EXISTS", key))
	if err != nil {
		UtilTools.LogError("判断键是否存在：" + err.Error())
		return false
	}
	return exists
}

//依据KEY获取REDIS数据
func GetValByKey(key string) string {
	client := global.RedisClient.Get()
	defer client.Close()
	obj, err := redis.String(client.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			return ""
		}
		UtilTools.LogError("获取REDIS数据失败：" + key + ";" + err.Error())
		return ""
	}
	return obj
}

//设置redis数据key-value
func SetRedisItem(key string, val string) {
	client := global.RedisClient.Get()
	defer client.Close()
	_, err := client.Do("SET", key, val)
	if err != nil {
		UtilTools.LogError("设置REDIS数据失败：" + key + ";" + val + err.Error())
	}
}

//设置redis数据key-value,并设置过期时间expiredTime，单位s
func SetRedisItemWithExpireTime(key string, val string, expiredTime int64) bool {
	client := global.RedisClient.Get()
	defer client.Close()
	_, err := client.Do("SET", key, val, "EX", strconv.FormatInt(expiredTime, 10))
	if err != nil {
		UtilTools.LogError("HASH设置REDIS数据失败：" + err.Error())
		return false
	}
	return true
}

//向REDIS队列中添加元素
func PushItemToList(list string, item string) bool {
	client := global.RedisClient.Get()
	defer client.Close()
	_, err := client.Do("RPUSH", list, item) //添加队列元素
	if err != nil {
		UtilTools.LogError("QUEYE PUSH数据失败：" + list + ";" + item + err.Error())
		return false
	}
	return true
}

//从REDIS LIST中删除指定元素
func RemoveItemFromList(list string, item string) bool {
	client := global.RedisClient.Get()
	defer client.Close()
	_, err := client.Do("LREM", list, 0, item) //删除队列中指定元素
	if err != nil {
		UtilTools.LogError("RemoveItemFromList REDIS数据失败：" + list + ";" + item + err.Error())
		return false
	}
	return true
}

//删除redis指定HASH项
func DelHashItem(hashKey string, key string) bool {
	client := global.RedisClient.Get()
	defer client.Close()
	_, err := client.Do("HDEL", hashKey, key)
	if err != nil {
		UtilTools.LogError("删除HASH项REDIS数据失败：" + hashKey + ":" + key + "：" + err.Error())
		return false
	}
	return true
}

//设置hash item
func SetHashItem(hashKey string, key string, value string) bool {
	client := global.RedisClient.Get()
	defer client.Close()
	_, err := client.Do("HSET", hashKey, key, value)
	if err != nil {
		UtilTools.LogError("HASH设置REDIS数据失败：" + hashKey + ":" + key + "：" + value + ";err:" + err.Error())
		return false
	}
	return true
}

//判断是否存在Hash Item
func ExistsHashItem(hashKey string, key string) bool {
	client := global.RedisClient.Get()
	defer client.Close()
	data, err := redis.Bool(client.Do("HEXISTS", hashKey, key))
	if err != nil {
		UtilTools.LogError("判断HASH ITEM是否存在，REDIS获取数据失败：" + hashKey + ":" + key + "err:" + err.Error())
		return false
	}
	return data
}

//获取HASHITEM
func GetHashItem(hashKey string, key string) string {
	client := global.RedisClient.Get()
	defer client.Close()
	data, err := redis.String(client.Do("HGET", hashKey, key))
	if err != nil {
		if err == redis.ErrNil { //不存在
			return ""
		}
		UtilTools.LogError("HASH REDIS获取数据失败：" + hashKey + ":" + key + "err:" + err.Error())
		return ""
	}
	return data
}
