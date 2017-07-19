package entities

import (
	"time"
)

//用户简要信息
type User struct {
	Id       int64
	UserName string
}

//用户请求票据
type UserTicket struct {
	UId         int64     //用户ID
	SigToken    string    //消息数字签名TOKEN
	SecretKey   string    //AES128加解密KEY
	ExpiredTime time.Time //票据过期时间
	CreateTime  time.Time //票据创建时间
}
