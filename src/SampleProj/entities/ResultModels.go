//统一返回消息格式
package entities

import (
	"commonlib/ffjson/ffjson"
	"net/http"
	"utiltools"
)

//统一返回结构数据
type ResultModel struct {
	StatusCode  int16  //返回状态码为HTTP标准状态码
	IsEncrypted bool   //返回前端提交是否对称加密
	Data        string //返回实际数据：加密JSON串，或是非加密JSON串
}

//返回加密后的JSON格式数据
func (result *ResultModel) ToEncryptJson(key []byte) string {
	jsonStr := utiltools.AesEncrypter([]byte(result.Data), key) //加密数据
	result.Data = jsonStr
	result.IsEncrypted = true //加密信息标示位
	return result.ToJson()
}

//非加密JSON化字符串
func (result *ResultModel) ToJson() string {
	json, err := ffjson.Marshal(result)
	if err != nil {
		utiltools.LogError("序列化JSON失败")
		return ""
	}
	return string(json)
}

//返回操作成功数据：data 序列化JSON
func OkResultModel(data string) *ResultModel {
	model := &ResultModel{
		StatusCode: http.StatusOK, //200
		Data:       data,
	}
	return model
}

//参数请求错误
func BadRequestResultModel(data string) *ResultModel {
	model := &ResultModel{
		StatusCode: http.StatusBadRequest,
		Data:       data,
	}
	return model
}

//内部处理错误
func InternalErrorResultModel(data string) *ResultModel {
	model := &ResultModel{
		StatusCode: http.StatusInternalServerError,
		Data:       data,
	}
	return model
}

//未找到数据错误
func NotFoundResultModel(data string) *ResultModel {
	model := &ResultModel{
		StatusCode: http.StatusNotFound,
		Data:       data,
	}
	return model
}

//未授权错误
func UnauthorizedResultModel(data string) *ResultModel {
	model := &ResultModel{
		StatusCode: http.StatusUnauthorized,
		Data:       data,
	}
	return model
}

//数据存在冲突：已经存在于数据库中
func ConflictResultModel(data string) *ResultModel {
	model := &ResultModel{
		StatusCode: http.StatusConflict,
		Data:       data,
	}
	return model
}

//TODO 可以依据需求再扩充其它类型的结果数据
