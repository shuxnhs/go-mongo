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
}
