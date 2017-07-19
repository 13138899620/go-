package global

//全局常量
const (
	DB_MYSQL_PROVIDER string = "mysql"

	DATE_FORMAT string = "2006-01-02"          //Date FORMAT
	TIME_FORMAT string = "2006-01-02 15:04:05" //TIME FORMAT

	//REQUEST HEADER 常量：UID,SIGNATURE,TIMESTAMP
	UID       string = "uid" //用户ID
	SIGNATURE string = "go-signature"
	TIMESTAMP string = "go-timestamp"

	POST_PARAM_NAME = "item" //统一POST参数名称

	//REQUEST HEADER NAME 登录COOKIE标识
	COOKIE_GO_USER string = "go-user"

	//获取临时TOKEN
	COOKIE_GO_TOKEN string = "go-token"
)
