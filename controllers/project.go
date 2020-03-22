package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"go-mongo/common/jwt"
	"go-mongo/models"
	"go-mongo/validate"
	"strings"
)

// 项目相关接口
type ProjectController struct {
	BaseController
}

// @Title 添加新的项目
// @Description 添加新的项目，并返回分配的mongo-key
// @Param	ProjectName		query 	string	true		"项目名"
// @Success 0 {string} 添加成功
// @Failure 1 {string} 添加失败
// @router /AddProject [get]
func (ctx *ProjectController) AddProject() {
	// 参数获取
	pn := ctx.GetString("ProjectName")

	// 参数校验
	v := &validate.ValidateProject{
		ProjectName: pn,
	}
	vErr := v.ValidAddProject()
	if vErr != nil {
		ctx.ApiError(400, fmt.Sprintf("%s", vErr))
	}
	project := models.Project{
		Project_name: pn,
		Is_deleted:   0,
	}
	// 添加新项目
	mongoKey, err := models.AddProject(project)
	if err != nil {
		ctx.ApiFail(1, fmt.Sprintf("%s", err))
	}
	ctx.ApiSuccessData("添加成功", mongoKey)
}

/**----------------后台接口------------------**/

// @Title: 获取项目全部配置
func (ctx *ProjectController) GetAllProject() {
	projects, err := models.GetAllProject()
	if err != nil {
		ctx.ApiFail(1, fmt.Sprintf("%s", err))
	}
	ctx.ApiSuccessData("获取成功", projects)
}

func (ctx *ProjectController) HandleLogin() {
	name := strings.TrimSpace(ctx.GetString("name"))
	password := strings.TrimSpace(ctx.GetString("password"))
	if name == "" || password == "" {
		ctx.TplName = "login.html"
		ctx.Data["errmsg"] = "请输入账号或密码"
		return
	}
	mdPassword := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	filter := models.FilterMap{
		"name":     name,
		"password": mdPassword,
		"role":     "admin",
	}
	mp := ctx.ApiMongoProxy()
	document, err := mp.FreeFindOne(filter)
	if err == nil && len(document) > 0 {
		// 登陆成功,生成jwt并设置cookie
		token, _ := jwt.CreateToken(name, mdPassword)
		ctx.Ctx.SetCookie("admin-jwt", token)
		ctx.TplName = "index.html"
	} else {
		ctx.TplName = "login.html"
		ctx.Data["errmsg"] = fmt.Sprintf("%s", err)
		return
	}
}

// @Title: 后台管理界面首页
func (ctx *ProjectController) Login() {
	ctx.TplName = "login.html"
}

// @Title: 后台管理界面首页
func (ctx *ProjectController) Index() {
	token := ctx.Ctx.GetCookie("admin-jwt")
	ctx.Ctx.Input.SetParam("mongoKey", "674D5122FEBC0F030C2AD55C9ED25B77")
	ctx.Ctx.Input.SetParam("collection", "user")
	if checkJwtToken(token, ctx) {
		ctx.TplName = "index.html"
	} else {
		ctx.TplName = "login.html"
	}
}

// @Title: 后台管理界面添加页
func (ctx *ProjectController) Add() {
	token := ctx.Ctx.GetCookie("admin-jwt")
	ctx.Ctx.Input.SetParam("mongoKey", "674D5122FEBC0F030C2AD55C9ED25B77")
	ctx.Ctx.Input.SetParam("collection", "user")
	if checkJwtToken(token, ctx) {
		ctx.TplName = "add.html"
	} else {
		ctx.TplName = "login.html"
	}
}

// @Title: 后台管理界面添加配置页
func (ctx *ProjectController) Config() {
	token := ctx.Ctx.GetCookie("admin-jwt")
	ctx.Ctx.Input.SetParam("mongoKey", "674D5122FEBC0F030C2AD55C9ED25B77")
	ctx.Ctx.Input.SetParam("collection", "user")
	if checkJwtToken(token, ctx) {
		ctx.TplName = "config.html"
	} else {
		ctx.TplName = "login.html"
	}
}

// 验证后台登陆的接口
func checkJwtToken(token string, ctx *ProjectController) bool {
	name, password, err := jwt.ParseToken(token)
	if err != nil {
		beego.Info(err)
		return false
	} else {
		beego.Info(name, password)
		mp := ctx.ApiMongoProxy()
		filter := models.FilterMap{
			"name":     name,
			"password": password,
			"role":     "admin",
		}
		beego.Info(name, filter)
		document, err := mp.FreeFindOne(filter)
		beego.Info(name, document)
		if err == nil && len(document) > 0 {
			return true
		} else {
			return false
		}
	}

}
