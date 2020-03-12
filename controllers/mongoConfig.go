package controllers

import (
	"fmt"
	"go-mongo/models"
	"go-mongo/validate"
)

// mongoDB配置相关接口
type MongoConfigController struct {
	BaseController
}

// @Title 添加mongodb配置
// @Description 添加配置接口
// @Param	mongoKey		query 	string	true		"项目mongo_key"
// @Param	host			query 	string	true		"项目mongodb的域名"
// @Param	port			query 	int	    true		"项目mongodb的端口"
// @Param	user			query 	string	false		"项目mongodb的用户"
// @Param	password		query 	string	false		"项目mongodb的密码"
// @Param   dbname			query   string	true		"项目mongodb的数据库名"
// @Success 200 {string} 添加成功
// @Failure 400 {string} 添加失败
// @router /AddMongoConfig [get]
func (ctx *MongoConfigController) AddMongoConfig() {
	// 参数获取
	mk := ctx.GetString("mongoKey")
	host := ctx.GetString("host")
	port, pErr := ctx.GetInt("port")
	if pErr != nil {
		port = 27017
	}
	user := ctx.GetString("user")
	dbname := ctx.GetString("dbname")
	pass := ctx.GetString("password")

	// 参数校验
	u := &validate.ValidateMongoConfig{
		MongoKey: mk,
		Host:     host,
	}
	vErr := u.ValidAddMongo()
	if vErr != nil {
		ctx.ApiError(400, fmt.Sprintf("%s", vErr))
	}

	// 判断是否存在
	if models.ExistMongoConfig(mk) {
		ctx.ApiFail(1, "配置已经存在")
	}

	// 添加新配置
	err := models.AddMongoConfigWith(mk, host, port, user, pass, dbname)
	if err != nil {
		ctx.ApiError(500, fmt.Sprintf("%s", err))
	}
	ctx.ApiSuccess("添加成功")
}

// @Title 保存mongodb配置
// @Description 添加配置接口
// @Param	mongoKey		query 	string	true		"项目mongo_key"
// @Param	host			query 	string	true		"项目mongodb的域名"
// @Param	port			query 	int		true		"项目mongodb的端口"
// @Param	user			query 	string	true		"项目mongodb的用户"
// @Param	password		query 	string	true		"项目mongodb的密码"
// @Param   dbname			query   string	true		"项目mongodb的数据库名"
// @Success 0 {string} 添加成功
// @Failure 1 {string} 添加失败
// @router /SaveMongoConfig [post]
func (ctx *MongoConfigController) SaveMongoConfig() {
	// 参数获取
	mk := ctx.GetString("mongoKey")
	host := ctx.GetString("host")
	port, pErr := ctx.GetInt("port")
	if pErr != nil {
		port = 27017
	}
	user := ctx.GetString("user")
	dbname := ctx.GetString("dbname")
	pass := ctx.GetString("password")

	// 判断是否存在
	config, err := models.GetMongoConfig(mk)
	ctx.ClearMkConnManger()
	if err != nil {
		// 添加新配置
		err := models.AddMongoConfigWith(mk, host, port, user, pass, dbname)
		if err != nil {
			ctx.ApiError(500, fmt.Sprintf("%s", err))
		}
	} else {
		// 更新配置
		err := models.UpdateMongoConfigWith(config.Id, host, port, user, pass, dbname)
		if err != nil {
			ctx.ApiError(500, fmt.Sprintf("%s", err))
		}
	}
	ctx.ApiSuccess("保存成功")
}

// @Title 删除mongodb配置
// @Description 删除mongodb配置
// @Param	mongoKey		query 	string	true		"项目mongo_key"
// @Success 200 {string} 删除成功
// @Failure 400 {string} 删除失败
// @router /DeleteMongoConfig [get]
func (ctx *MongoConfigController) DeleteMongoConfig() {
	// 参数获取
	ctx.ApiCheckCommon()

	// 进行操作
	mk := ctx.ApiGetMongoKey()
	delRs, err := models.DeleteMongoConfig(mk)
	if err != nil {
		ctx.ApiError(500, fmt.Sprintf("%s", err))
	}
	if !delRs {
		ctx.ApiFail(1, "配置不存在")
	}
	ctx.ClearMkConnManger()
	ctx.ApiSuccess("删除成功")
}

// @Title 获取mongodb配置
// @Description 获取mongodb配置
// @Param	mongoKey		query 	string	true		"项目mongo_key"
// @Success 0 {string}	配置存在
// @Failure 1 {string}  配置不存在
// @router /GetMongoConfig [get]
func (ctx *MongoConfigController) GetMongoConfig() {
	ctx.ApiCheckCommon()

	// 获取配置
	mk := ctx.ApiGetMongoKey()
	config, err := models.GetMongoConfig(mk)
	if err != nil {
		ctx.ApiFail(1, "配置不存在")
	}

	// 隐藏密码
	config.Password = "******"
	ctx.ApiSuccessData("获取成功", config)
}

// @Title 检测是否可以连接上mongoDB
// @Description 检测是否可以连接上mongoDB
// @Param	mongoKey		query 	string	true		"项目mongo_key"
// @Success 0 {string}	连接成功
// @Failure 1 {string}  连接失败
// @router /CheckMongoConnect [get]
func (ctx *MongoConfigController) CheckMongoConnect() {
	ctx.ApiCheckCommon()
	ctx.Input().Add("collection", "test") // ApiMongoProxy需要collection
	mp := ctx.ApiMongoProxy()
	err := mp.CheckPing()
	if err != nil {
		ctx.ApiFail(1, "连接失败")
	}
	ctx.ApiSuccess("连接成功")
}
