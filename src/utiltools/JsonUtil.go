package utiltools

import "commonlib/ffjson/ffjson"

//序列化
func ToJson(obj interface{}) string {
	json, err := ffjson.Marshal(obj)
	if err != nil {
		LogError("序列化失败" + err.Error())
		return ""
	}
	return string(json)
}

//返回序列化为对象
func FromJson(json string, obj interface{}) error {
	err := ffjson.Unmarshal([]byte(json), obj)
	if err != nil {
		return err
	}
	return nil
}

//返回序列化对象
func FromByteJson(json []byte, obj interface{}) error {
	err := ffjson.Unmarshal(json, obj)
	if err != nil {
		return err
	}
	return nil
}
