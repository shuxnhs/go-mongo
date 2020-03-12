package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go.mongodb.org/mongo-driver/mongo/options"
	// 需要使用mysql驱动来初始化
	_ "github.com/go-sql-driver/mysql"
	utiLog "go-mongo/common/log"
)

var clientOptions *options.ClientOptions

func init() {

	// 在此统一初始化
	dbHost := beego.AppConfig.String("dbHost")
	dbPort := beego.AppConfig.String("dbPort")
	dbUser := beego.AppConfig.String("dbUser")
	dbPass := beego.AppConfig.String("dbPass")
	dbName := beego.AppConfig.String("dbName")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Asia%2FShanghai"

	orm.RegisterModel(new(Config), new(Project))
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err == nil {
		utiLog.Log.Info("数据库连接成功️")
	} else {
		utiLog.Log.Error("数据库连接失败: ", err)
		beego.Info("数据库连接失败,失败原因：", err)
	}

	// 初始化mongoClient
	clientOptions = options.Client().
		SetMaxPoolSize(1000).       // 连接池中每个服务器允许的最大连接数
		SetMaxConnIdleTime(5 * 60). // 在连接池中保持空闲状态的最长时间
		//SetConnectTimeout(3).		// 与服务器的连接的超时时间
		//SetDialer(&net.Dialer{KeepAlive: 60 * 60,}).
		SetRetryReads(true). // 针对某些错误（例如网络错误）重试一次支持的读取操作
		SetRetryWrites(true) // 针对某些错误（例如网络错误）重试一次支持的写入操作
}
