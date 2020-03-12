package main

import (
	"github.com/astaxie/beego"
	_ "go-mongo/common/log" // 日志模块加载
	"go-mongo/controllers"
	"go-mongo/models"
	_ "go-mongo/routers"
	"go-mongo/validate"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.EnableErrorsShow = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		// 开启热升级
		beego.BConfig.Listen.Graceful = true
	}

	/*-------------验证信息的初始化定义--------------*/
	validate.InitValidate()

	/*--------------注册全局mongoProxy-------------*/
	controllers.MongoProxy = models.NewMongoProxy()

	beego.Run()
}
