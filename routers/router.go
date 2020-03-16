// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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
	)
	beego.AddNamespace(ns)
}
