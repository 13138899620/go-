package facade

import (
	Entities "SampleProj/entities"
	UserEntities "SampleProj/entities/user"
	UserModel "SampleProj/models/user"
)

//获取用户所有信息列表
func GetUserList() *Entities.ResultModel {
	return UserModel.GetUserList()
}

//添加用户信息
func AddUser(user *UserEntities.User) *Entities.ResultModel {
	return UserModel.AddUser(user)
}
