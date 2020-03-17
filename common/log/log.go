package log

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-mongo/common/time"
	"strconv"
)

var Log *logs.BeeLogger

func init() {
	Log = logs.NewLogger(1000)

	/*
		filename 保存的文件名
		maxlines 每个文件保存的最大行数，默认值 1000000
		maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
		daily 是否按照每天 logrotate，默认是 true
		maxdays 文件最多保存多少天，默认保存 7 天
		rotate 是否开启 logrotate，默认是 true
		level 日志保存的时候的级别，默认是 Trace 级别
		perm 日志文件权限
	*/

	logDir := beego.AppConfig.String("fileLog") // 日志文件目录
	date := time.GetDate
	logFile := logDir + "go-mongo" + date() + ".log" // log文件路径和名称
	level := strconv.Itoa(logs.LevelInfo)            // 日志级别
	maxsize := "10240"                               // 日志文件大小， 上线时请设置为0

	// 设置输入log的方式
	err := Log.SetLogger("file", `{"filename":"`+logFile+`","level":`+level+`,"maxlines":0,"maxsize":`+maxsize+`,"maxdays":0}`)

	if err != nil {
		panic("设置日志配置信息失败, err: " + err.Error())
	}

	// 日志双写到ES
	// 官方esAdapter有问题，注册自己的esAdapter
	isEnableEsLog, err := beego.AppConfig.Bool("EnableEsLog")
	if err == nil && isEnableEsLog {
		logs.Register("esLog", NewES)
		Log = logs.NewLogger(1000)
		dsn := beego.AppConfig.String("esLog")
		// 将log同步到es
		err2 := Log.SetLogger("esLog", `{"dsn":"`+dsn+`"}`)
		if err2 != nil {
			panic("日志输出到Es失败，err： " + err2.Error())
		}
	}

	// 输出调用的文件名和文件行号
	Log.EnableFuncCallDepth(true)

	// 异步输出
	Log.Async()

}
