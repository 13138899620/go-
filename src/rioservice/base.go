package rioservice

import (
	"bytes"
	"commonlib/ffjson/ffjson"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"global"
	"io/ioutil"
	"net/http"
	"time"
)

//构建邮件
func buildMail(sender, receiver, title, content string) (ms []byte, err error) {
	source := fmt.Sprintf("appkey:%s,name:%s", global.ConfigMappings[global.TOF_APP_KEY], global.ConfigMappings[global.SYSTEM_NAME])
	//构造邮件内容
	mailBody := MailBody{
		EmailType:     InternalEmail,
		Priority:      NormalPriority,
		BodyFormat:    FormatHtml,
		MessageStatus: StatusWait,
		From:          sender,
		To:            receiver,
		Title:         title,
		Content:       content,
		StartTime:     time.Now(),
		EndTime:       time.Now().AddDate(0, 0, 7),
	}
	mail := MSMail{
		Mail:   mailBody,
		Source: source,
	}
	//检验
	if err = mail.Validate(); err != nil {
		return
	}
	//json化
	ms, err = ffjson.Marshal(mail)
	return
}

func postData(url string, data []byte) (err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", url, bytes.NewBuffer(data)); err != nil {
		err = fmt.Errorf("[NewRequest] %v", err)
		return
	}
	// 3. 设置头信息
	req.Header.Set("Content-Type", "application/json")
	// 4. 建立Client，发起请求
	client := http.Client{}
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		err = fmt.Errorf("[Client] %v", err)
		return
	}
	defer resp.Body.Close()
	// 5. 获取响应内容
	var respbody []byte
	if respbody, err = ioutil.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("[ReadAll] %+v", err)
		return
	}
	// 5.1 如果状态码不为200，报错
	if resp.StatusCode != 200 {
		err = fmt.Errorf("[respond state code] code:%d, body:%s", resp.StatusCode, string(respbody))
		return
	}
	// 5.2 如果对象序列化失败，报错
	respond := RespondModel{}
	if err = json.Unmarshal(respbody, &respond); err != nil {
		err = fmt.Errorf("[unmarshal resp body] %v", err)
		return
	}

	// 5.3 如果返回的 Ret 不为0，失败
	// if respond.SendSMSResult != nil {
	// 	err = fmt.Errorf("[respond state ret] %v", respond.SendSMSResult)
	// 	return
	// }
	return
}

//构建创建myoa的请求内容
func buildMyoaCreateBody(guid, formUrl, handler, startTime, dueTime, periodName, objName, roleName string) (data []byte, err error) {

	var detailViews []MyoaDetailView

	detailView1 := MyoaDetailView{
		Key:   "周期名称",
		Value: periodName,
	}
	detailView2 := MyoaDetailView{
		Key:   "您的角色",
		Value: roleName,
	}
	detailView3 := MyoaDetailView{
		Key:   "被评估对象",
		Value: objName,
	}
	detailViews = append(detailViews, detailView1, detailView2, detailView3)

	body := MyoaCreateBody{
		Category:      global.ConfigMappings[global.MYOA_CATEGORY],
		ProcessName:   "360评估系统",
		ProcessInstId: "360评估系统-" + guid, //AssessorGuid 或AssessorId 为唯一标识
		Activity:      "360评估系统待办",
		Title:         "360评估系统:" + periodName + ",评估对象:" + objName + ",您的角色:" + roleName,
		FormUrl:       formUrl, //业务单据的访问地址
		MobileFormUrl: formUrl, //业务单据的移动端访问地址,移动端和pc端是相同的
		CallbackUrl:   formUrl, //回调地址，暂不需要
		// 去掉与审批相关的参数，会返回400
		EnableQuickApproval: true,
		EnableBatchApproval: true,
		Handler:             handler, //当前审批单据的处理人的英文名
		//
		Applicant:   "360admin", //申请人，其实我不知道是谁
		StartTime:   startTime,  //RFC3339格式，"2016-06-29T11:10:52.0Z"
		DueTime:     dueTime,
		DetailViews: detailViews, //需要显示的基本信息
	}
	var items MyoaCreateItems
	items.WorkItems = append(items.WorkItems, body)

	data, err = ffjson.Marshal(items)
	if err != nil {
		fmt.Println("json err:", err)
	}

	//可算是把数据组织好了
	fmt.Println(string(data))
	return
}

//构建撤回myoa的请求内容
func buildMyoaDelBody(guid, handler string) (data []byte, err error) {

	body := MyoaCloseBody{
		Category:      global.ConfigMappings[global.MYOA_CATEGORY],
		ProcessName:   "360评估系统",
		ProcessInstId: "360评估系统-" + guid, //AssessorGuid 或AssessorId 为唯一标识
		Activity:      "360评估系统撤回待办",
		Handler:       handler, //当前审批单据的处理人的英文名
	}

	data, err = ffjson.Marshal(body)
	if err != nil {
		fmt.Println("json err:", err)
	}

	//可算是把数据组织好了
	fmt.Println(string(data))
	return
}

//myoa
func postMyoaData(url string, data []byte) (err error) {
	fmt.Println(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	sn := timestamp + global.ConfigMappings[global.MYOA_TOKEN] + timestamp
	signature := fmt.Sprintf("%x", sha256.Sum256([]byte(sn)))

	req.Header.Set("signature", signature)
	req.Header.Set("timestamp", timestamp)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		err = fmt.Errorf("[Client] %v", err)
		return
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	// 5. 获取响应内容
	var respbody []byte
	if respbody, err = ioutil.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("[ReadAll] %+v", err)
		return
	}
	// 5.1 如果状态码不为200，报错
	if resp.StatusCode != 200 {
		err = fmt.Errorf("[respond state code] code:%d, body:%s", resp.StatusCode, string(respbody))
		return
	}
	// 5.2 如果对象序列化失败，报错
	respond := RespondModel{}
	if err = json.Unmarshal(respbody, &respond); err != nil {
		err = fmt.Errorf("[unmarshal resp body] %v", err)
		return
	}
	return
}
