package rioservice

import (
	"fmt"
	"strings"
	"time"
)

const (
	ExternalEmail = 0 //外部邮件
	InternalEmail = 1 //内部邮件
	MeetingEmail  = 2 //会议邮件

	LowPriority    = -1 //低优先级
	NormalPriority = 0  //普通优先级
	HighPriority   = 1  //高优先级

	FormatTxt  = 0 //文本格式
	FormatHtml = 1 //Html格式

	StatusWait = 0 //等待发送
)

type MSMail struct {
	Mail   MailBody `json:"mail"`
	Source string   `json:"source"`
}

type MailBody struct {
	EmailType     int
	From          string
	To            string
	CC            string
	Bcc           string
	Content       string
	Title         string
	Priority      int
	BodyFormat    int
	Location      string
	Organizer     string
	StartTime     time.Time
	EndTime       time.Time
	Attachment    []AttachmentBody
	MessageStatus int
}

type AttachmentBody struct {
	FileName    string
	FileContent []byte
}

type RespondModel struct {
	SendSMSResult string
}

func (el *MSMail) Validate() (err error) {
	if strings.TrimSpace(el.Mail.From) == "" {
		err = fmt.Errorf("Sender is empty")
		return
	}
	if strings.TrimSpace(el.Mail.To) == "" {
		err = fmt.Errorf("Receiver is empty")
		return
	}
	if strings.TrimSpace(el.Mail.Content) == "" {
		err = fmt.Errorf("Content is empty")
	}
	if strings.TrimSpace(el.Mail.Title) == "" {
		err = fmt.Errorf("Title is empty")
	}
	return
}

type MyoaCloseBody struct {
	Category      string `json:"category"`
	ProcessName   string `json:"process_name"`
	ProcessInstId string `json:"process_inst_id"`
	Activity      string `json:"activity"`
	Handler       string `json:"handler"`
}

type MyoaCreateItems struct {
	WorkItems []MyoaCreateBody `json:"work_items"`
}

type MyoaCreateBody struct {
	Category            string           `json:"category"`
	ProcessName         string           `json:"process_name"`
	ProcessInstId       string           `json:"process_inst_id"`
	Activity            string           `json:"activity"`
	Title               string           `json:"title"`
	FormUrl             string           `json:"form_url"`
	MobileFormUrl       string           `json:"mobile_form_url"`
	CallbackUrl         string           `json:"callback_url"`
	EnableQuickApproval bool             `json:"enable_quick_approval"`
	EnableBatchApproval bool             `json:"enable_batch_approval"`
	Handler             string           `json:"handler"`
	Applicant           string           `json:"applicant"`
	StartTime           string           `json:"start_time"`
	DueTime             string           `json:"due_time"`
	DetailViews         []MyoaDetailView `json:"detail_view"`
}

//myoa待办基本信息
type MyoaDetailView struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
