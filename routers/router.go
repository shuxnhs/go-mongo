// @APIVersion 1.0.0
// @Title Go-Mongo
// @Description go的mongodb中间件，为项目提供mongodb的API
// @Contact 610087273@qq.com
// @TermsOfServiceUrl http://beego.me/
package routers

import (
	"github.com/astaxie/beego"
	"go-mongo/controllers"
)

func init() {
	// 对外开放的接口
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),
		beego.NSNamespace("/mongodb",
			beego.NSInclude(
				&controllers.MongoConfigController{},
			),
		),
		beego.NSNamespace("/mongoC",
			beego.NSInclude(
				&controllers.MongoCController{},
			),
		),
		beego.NSNamespace("/mongoU",
			beego.NSInclude(
				&controllers.MongoUController{},
			),
		),
		beego.NSNamespace("/mongoR",
			beego.NSInclude(
				&controllers.MongoRController{},
			),
		),
		beego.NSNamespace("/mongoD",
			beego.NSInclude(
				&controllers.MongoDController{},
			),
		),
		beego.NSNamespace("/mongoI",
			beego.NSInclude(
				&controllers.MongoIController{},
			),
		),
		beego.NSNamespace("/mongoLBS",
			beego.NSInclude(
				&controllers.MongoLBSController{},
			),
		),
		beego.NSNamespace("/mongoFT",
			beego.NSInclude(
				&controllers.MongoFTController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// 管理后台界面路由
	beego.Router("/login", &controllers.AdminController{}, "*:Login")
	beego.Router("/admin", &controllers.AdminController{}, "*:Index")
	beego.Router("/add", &controllers.AdminController{}, "*:Add")
	beego.Router("/config", &controllers.AdminController{}, "*:Config")
	beego.Router("/handleLogin", &controllers.AdminController{}, "*:HandleLogin")

	// 不使用路由注解的接口,不生成接口文档，不对外开放
	beego.Router("/project/getAllProject", &controllers.ProjectController{}, "*:GetAllProject")

	// 静态资源加载
	beego.SetStaticPath("/views", "views")
}
