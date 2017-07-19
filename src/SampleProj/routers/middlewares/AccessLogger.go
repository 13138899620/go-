// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routers

import (
	gin "commonlib/gin"
	"fmt"
	"global"
	"time"
	"utiltools"
)

//访问日志记录
func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		c.Next() // Process request
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.String()
		//TODO 获取访问用户ID信息，可能是空的,用户可以自定义
		uid := c.Request.Header.Get(global.UID)
		//采用自定义写文件方式
		content := fmt.Sprintf("%s,%3d,%13v,%s,%s,%-7s,%s,%s",
			end.Format("2006/01/02 15:04:05"),
			statusCode,
			latency,
			clientIP,
			uid,
			method,
			path,
			comment)
		utiltools.LogInfo(content) //写访问日志
	}
}
