package models

import (
	entities "SampleProj/entities"
	UserEntities "SampleProj/entities/user"
	DBCommon "SampleProj/models"
	"commonlib/ffjson/ffjson"
	"strconv"
	"utiltools"
)

//获取用户列表
func GetUserList() *entities.ResultModel {
	db, err := DBCommon.InitDbConn()
	if err != nil {
		//TODO 记录日志发邮件
		utiltools.LogError(err.Error())
		return entities.InternalErrorResultModel("数据获取失败")
	}
	defer db.Close()
	sql := `SELECT 
    Id, UserName
FROM
    user
WHERE
    EnableFlag = TRUE`
	rows, rowsErr := db.Query(sql)
	if rowsErr != nil {
		//TODO 记录日志发邮件
		utiltools.LogError(rowsErr.Error())
		return entities.InternalErrorResultModel("数据获取失败")
	}
	defer rows.Close()
	result := []UserEntities.User{}
	for rows.Next() {
		tmp := UserEntities.User{}
		rows.Scan(&tmp.Id, &tmp.UserName)
		result = append(result, tmp)
	}
	//返回结果数据
	json, _ := ffjson.Marshal(result)
	return entities.OkResultModel(string(json))
}

//新增用户
func AddUser(user *UserEntities.User) *entities.ResultModel {
	db, err := DBCommon.InitDbConn()
	if err != nil {
		//TODO 记录日志发邮件
		utiltools.LogError(err.Error())
		return entities.InternalErrorResultModel("添加数据失败")
	}
	defer db.Close()
	stmtIns, stErr := db.Prepare("INSERT INTO user(UserName,EnableFlag) VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		utiltools.LogError(stErr.Error())
		return entities.InternalErrorResultModel("添加数据失败")
	}
	defer stmtIns.Close()
	sqlResult, rErr := stmtIns.Exec(user.UserName, 1)
	if rErr != nil {
		utiltools.LogError(rErr.Error())
		return entities.InternalErrorResultModel("添加数据失败")
	}
	id, _ := sqlResult.LastInsertId()
	return entities.OkResultModel(strconv.FormatInt(id, 10)) //return ID
}
