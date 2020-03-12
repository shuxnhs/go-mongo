package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-mongo/common/apiResponse"
	utiLog "go-mongo/common/log"
	"go-mongo/models"
	"go-mongo/validate"
)

/**
 * API基类
 */
type BaseController struct {
	beego.Controller
}

var MongoProxy *models.MongoProxy

/**
 * 异常返回
 */
func (ctx *BaseController) ApiError(ret int, msg string) {
	res := apiResponse.NewAPIResponse()
	data := apiResponse.NewAPIDataResponse()
	ctx.Data["json"] = res.SetData(data).Error(ret, msg)
	ctx.ServeJSON()
	utiLog.Log.Error("异常返回：", res, ctx.getApiInfoToLog())
	ctx.StopRun()
}

/**
 * 异常返回，带数据
 */
func (ctx *BaseController) ApiErrorData(ret int, msg string, data interface{}) {
	res := apiResponse.NewAPIResponse()
	resData := apiResponse.NewAPIDataResponse()
	resData.SetData(data)
	ctx.Data["json"] = res.SetData(resData).Error(ret, msg)
	ctx.ServeJSON()
	utiLog.Log.Error("异常返回：", res, ctx.getApiInfoToLog())
	ctx.StopRun()
}

/**
 * 业务失败返回
 */
func (ctx *BaseController) ApiFail(code int, msg string) {
	res := apiResponse.NewAPIResponse()
	data := apiResponse.NewAPIDataResponse()
	data.SetCode(code).SetMsg(msg)
	ctx.Data["json"] = res.SetData(data)
	ctx.ServeJSON()
	utiLog.Log.Error("业务失败返回：", res, ctx.getApiInfoToLog())
	ctx.StopRun()
}

/**
 * 业务失败返回，带数据
 */
func (ctx *BaseController) ApiFailData(code int, msg string, data interface{}) {
	res := apiResponse.NewAPIResponse()
	resData := apiResponse.NewAPIDataResponse()
	resData.SetCode(code).SetMsg(msg).SetData(data)
	ctx.Data["json"] = res.SetData(resData)
	ctx.ServeJSON()
	utiLog.Log.Error("业务失败返回：", res, ctx.getApiInfoToLog())
	ctx.StopRun()
}

/**
 * 业务成功返回
 */
func (ctx *BaseController) ApiSuccess(msg string) {
	res := apiResponse.NewAPIResponse()
	data := apiResponse.NewAPIDataResponse()
	data.SetMsg(msg)
	ctx.Data["json"] = res.SetData(data)
	ctx.ServeJSON()
	utiLog.Log.Error("业务成功返回：", ctx.GetControllerAndAction, res, ctx.getApiInfoToLog())
	ctx.StopRun()
}

/**
 * 业务成功返回，带数据
 */
func (ctx *BaseController) ApiSuccessData(msg string, data interface{}) {
	res := apiResponse.NewAPIResponse()
	resData := apiResponse.NewAPIDataResponse()
	resData.SetMsg(msg).SetData(data)
	ctx.Data["json"] = res.SetData(resData)
	ctx.ServeJSON()
	utiLog.Log.Error("业务成功返回：", res, ctx.getApiInfoToLog())
	ctx.StopRun()
}

func (ctx *BaseController) getApiInfoToLog() string {
	con, act := ctx.GetControllerAndAction()
	return "controller: " + con + ", action: " + act
}

/**
 * 参数检测
 */
func (ctx *BaseController) ApiCheckCommon() {
	u := &validate.ValidateMongoKey{
		Mongo_key: ctx.GetString("mongoKey"),
	}
	vErr := u.ValidateOperater()
	if vErr != nil {
		ctx.ApiError(400, fmt.Sprintf("%s", vErr))
		utiLog.Log.Error("参数检测失败：：", fmt.Sprintf("%s", vErr))
		ctx.StopRun()
	}
}

func (ctx *BaseController) ApiGetMongoKey() string {
	return ctx.GetString("mongoKey")
}

/**
 * 获取Mongodb
 */
func (ctx *BaseController) ApiMongoProxy() *models.MongoConnection {
	ctx.ApiCheckCommon()
	mk := ctx.ApiGetMongoKey()
	if !models.ExistMongoConfig(mk) {
		ctx.ApiError(400, "mongo-key不存在")
	}
	collection := ctx.GetString("collection")
	if collection == "" {
		ctx.ApiError(400, "collection不能为空")
	}
	conn, err := MongoProxy.GetMongoProxyManager().GetConn(mk)
	if err != nil {
		conn := MongoProxy.Start(mk, collection)
		if conn == nil {
			ctx.ApiError(400, "配置错误")
		}
		return conn
	}
	conn.SwitchCollection(collection)
	return conn
}

/**
 * 更新config配置删掉连接池中的连接
 */
func (ctx *BaseController) ClearMkConnManger() {
	mk := ctx.ApiGetMongoKey()
	// 删除连接池中的配置
	conn, err := MongoProxy.MongoConnectionManager.GetConn(mk)
	if err == nil {
		MongoProxy.MongoConnectionManager.DelConn(conn)
	}
}

func (ctx *BaseController) ClearAllConnManger() {
	MongoProxy.MongoConnectionManager.DelAllConn()
}

func (ctx *BaseController) GetConnNum() int {
	return MongoProxy.MongoConnectionManager.GetConnNum()
}
