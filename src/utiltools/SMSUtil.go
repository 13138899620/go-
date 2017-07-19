/*腾讯云短信服务发送接口*/
package utiltools

import (
	"crypto/md5"
	"fmt"
	"global"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//短信发送的地址
var postUrl string = "https://yun.tim.qq.com/v3/tlssmssvr/sendsms?sdkappid=%s&random=%s"

//消息结构体
var msgbody string = `{
	"tel":{
		"nationcode":"86",
		"phone":"%s"
	},
	"type":"0",
	"msg":"%s",
	"sig":"%s",
	"extend":"",
	"ext":""
}`

//消息模板：这个不能随便改，要与腾讯云上面的消息模板匹配才行
var msgTmpl string = "【云评估】您还有360度评估问卷未提交，请按时提交:%s"

//格式化消息体
func getMsgBody(mobile string, content string, random string) string {
	str := global.ConfigMappings[global.MSG_App_Key] + mobile
	signature := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return fmt.Sprintf(msgbody, mobile, content, signature)
}

//发送后返回的结构体
type MsgSendResult struct {
	result string //0 表示成功，非0表示失败
	errmsg string //非0时的具体错误信息
	ext    string //用户的SESSION内容
	sid    string //标识本次发送的ID
	fee    string //短信计费的条数
}

//发送短信API
//shortUrl为模板消息里面待替换内容
func SendSMSMsg(mobile string, shortUrl string) *MsgSendResult {
	seed := int64(time.Now().UTC().Nanosecond())
	rand.Seed(seed)
	random := strconv.FormatInt(rand.Int63n(100000)%(900000)+100000, 10)
	content := fmt.Sprintf(msgTmpl, shortUrl) //短信内容
	body := getMsgBody(mobile, content, random)
	msgPostUrl := fmt.Sprintf(postUrl, global.ConfigMappings[global.MSG_App_ID], random)
	result := postMsg(body, msgPostUrl)
	return result
}

//发送数据
func postMsg(json string, url string) *MsgSendResult {
	body := strings.NewReader(json)
	header := "application/x-www-form-urlencoded;charset=utf-8"
	res, err := http.Post(url, header, body)
	if err != nil {
		//创建短信发送链接失败
		LogWarn("短信发送失败" + err.Error() + ";" + json + ";" + url)
		return nil
	}
	result, rerr := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if rerr != nil {
		//创建短信发送链接失败
		LogWarn("短信发送读取返回结果失败" + rerr.Error() + ";" + json + ";" + url)
		return nil
	}
	var returnResult MsgSendResult
	FromJson(string(result), &returnResult) //ffjson 反序列化为对象
	if returnResult.result != "0" {
		LogWarn("短信发送失败;" + json + ";" + url + ";return result:" + string(result))
		return nil
	}
	LogWarn("短信发送成功：" + string(result))
	return &returnResult
}
