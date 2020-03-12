package controllers

import (
	"encoding/json"
	"fmt"
	"go-mongo/models"
)

// mongo文档读取操作相关接口
type MongoRController struct {
	BaseController
}

// @Title 获取文档的条数
// @Description 自由获取文档的条数, 返回：{ "ret": 200, "msg": "", "data": { "code": 0, "msg": "获取成功", "data": 4}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   filter  		query	string  true		"查找文档的条件，查找全部传空json，如{"name":"hxh"}"
// @Success 200 {string} 更新成功的文档数量/更新失败的错误信息
// @router /CountData [get]
func (ctx *MongoRController) CountData() {
	filter := ctx.GetString("filter")
	var filterWhere models.FilterMap
	err1 := json.Unmarshal([]byte(filter), &filterWhere)
	if err1 != nil {
		ctx.ApiError(400, "filter格式错误")
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.CountData(filterWhere)

	if err != nil {
		ctx.ApiFailData(1, "查询失败", result)
	}
	ctx.ApiSuccessData("获取成功", result)
}

// @Title 获取文档
// @Description 获取collection指定文档
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   objectId		query	string   true		"对象ID"
// @Success 0 {string} 清空成功
// @Failure 1 {string} 清空失败错误信息
// @router /Retrieve [get]
func (ctx *MongoRController) Retrieve() {
	objectId := ctx.GetString("objectId")

	mp := ctx.ApiMongoProxy()
	object, err := mp.RetrieveObject(objectId)

	if err != nil {
		ctx.ApiFail(1, fmt.Sprintf("%s", err))
	}
	ctx.ApiSuccessData("获取成功", object)
}

// @Title 获取n份文档
// @Description 获取collection指定条数文档{"ret": 200, "msg": "", "data": {"code": 0, "msg": "查询成功", "data": [{"_id": "5da368b34a0bab8c655a8142", "age": 18, "name": "hxh"}]}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   num				query	int64   true		"获取条数"
// @Success 0 {string} 获取成功
// @Failure 1 {string} 集合文档为空
// @Failure 2 {string} mongoDb其他查询错误
// @router /GetDataList [get]
func (ctx *MongoRController) GetDataList() {
	num, err1 := ctx.GetInt64("num")
	if err1 != nil || num < 1 {
		ctx.ApiError(400, "num参数错误")
	}
	mp := ctx.ApiMongoProxy()
	result, err := mp.GetDataList(num)

	if err != nil {
		ctx.ApiFailData(2, fmt.Sprintf("%s", err), result)
	}

	if len(result) == 0 {
		ctx.ApiFail(1, "集合文档为空")
	}
	ctx.ApiSuccessData("查询成功", result)
}

// @Title 自由获取文档
// @Description 通过条件，获取一条文档,{"ret": 200, "msg": "", "data": {"code": 0, "msg": "查询成功", "data": {"_id": "5da368b34a0bab8c655a8142", "age": 18, "name": "hxh"}}}
// @Param	mongoKey		query 	string	true		"用户mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   filter          query   string  true        "查找文档的条件，json传递，如{"name":"hxh"}"
// @router /FreeFindOne [get]
// @Success 0 {string} 获取成功
// @Failure 1 {string} 集合文档为空
func (ctx *MongoRController) FreeFindOne() {
	filter := ctx.GetString("filter")
	var filterWhere models.FilterMap
	err := json.Unmarshal([]byte(filter), &filterWhere)
	if err != nil {
		ctx.ApiError(400, "filter 格式错误")
	}
	mp := ctx.ApiMongoProxy()
	object, fErr := mp.FreeFindOne(filterWhere)

	if fErr != nil {
		ctx.ApiFail(1, "集合文档为空")
	}
	ctx.ApiSuccessData("查询成功", object)
}

// @Title 自由获取文档
// @Description 自由获取collection指定条数文档,{"ret": 200, "msg": "", "data": {"code": 0, "msg": "查询成功", "data": {"_id": "5da368b34a0bab8c655a8142", "age": 18, "name": "hxh"}}}
// @Param	mongoKey		query 	string	true		"用户mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   projection  	query	string  false		"指定返回的字段，比如只返回name就{"name":1},不返回name就{"name":0},默认返回全部"
// @Param   filter  		query	string  true		"条件查找文档的条件，如{"name":"hxh"}"
// @Param   num				query	int64   true		"获取条数"
// @Success 0 {string} 清空成功
// @Failure 1 {string} 清空失败错误信息
// @router /FreeGetDataList [get]
func (ctx *MongoRController) FreeGetDataList() {
	num, err1 := ctx.GetInt64("num")
	if err1 != nil || num < 1 {
		ctx.ApiError(400, "num参数错误")
	}

	// 过滤条件
	filter := ctx.GetString("filter")
	var filterWhere models.FilterMap
	err2 := json.Unmarshal([]byte(filter), &filterWhere)
	if err2 != nil {
		ctx.ApiError(400, "filter格式错误")
	}

	// 返回的字段
	projection := ctx.GetString("projection")
	var projectionOpt models.ProjectionMap
	if projection != "" {
		err3 := json.Unmarshal([]byte(projection), &projectionOpt)
		if err3 != nil {
			ctx.ApiError(400, "projection格式错误")
		}
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.FreeGetDataList(projectionOpt, filterWhere, num)

	if err != nil {
		ctx.ApiFail(1, "查询失败"+fmt.Sprintf("%s", err))
	}

	ctx.ApiSuccessData("获取成功", result)
}
