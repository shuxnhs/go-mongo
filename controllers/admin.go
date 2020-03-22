package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"go-mongo/common/jwt"
	"go-mongo/models"
	"strings"
)

// 管理后台模块
type AdminController struct {
	BaseController
}

func (ctx *AdminController) HandleLogin() {
	ctx.Ctx.Input.SetParam("mongoKey", "674D5122FEBC0F030C2AD55C9ED25B77")
	ctx.Ctx.Input.SetParam("collection", "user")
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
func (ctx *AdminController) Login() {
	ctx.TplName = "login.html"
}

// @Title: 后台管理界面首页
func (ctx *AdminController) Index() {
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
func (ctx *AdminController) Add() {
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
func (ctx *AdminController) Config() {
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
func checkJwtToken(token string, ctx *AdminController) bool {
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
