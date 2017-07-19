//数据库公共操作
//建议直接采用mysql库
//https://github.com/go-sql-driver/mysql
//使用方法：https://github.com/go-sql-driver/mysql/wiki/Examples
//虽然没有.NET写起来那么顺手，但基本够用，也可以借此机会多练练SQL基础语法
//如果真的想用orm组件，推荐用GORM,BEEGO(这个比较重)
package models

import (
	_ "commonlib/mysqldriver"
	"database/sql"
	"global"
)

//初化链接
func InitDbConn() (db *sql.DB, err error) {
	dbConStr := global.ConfigMappings[global.DB_CONN]
	//这里采用mysql驱动
	dbConn, err := sql.Open(global.DB_MYSQL_PROVIDER, dbConStr)
	return dbConn, err
}
