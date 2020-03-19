package controllers

import (
	"encoding/json"
	"fmt"
	"go-mongo/models"
)

// mongo文档更新操作相关接口
type MongoUController struct {
	BaseController
}

// @Title 根据_id更新文档
// @Description 根据_id更新文档,{"ret": 200, "msg": "", "data": {"code": 0, "msg": "更新成功", "data": {"MatchedCount": 1, "ModifiedCount": 1, "UpsertedCount": 0, "UpsertedID": null}}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   objectId		query	string   true		"对象ID"
// @Param   update  		query	string  true		"条件更新文档的数据，json传递，如{"name":"hxh"}"
// @Success 0 {string} 获取成功
// @Failure 1 {string} 获取失败错误信息
// @router /UpdateDataById [post]
func (ctx *MongoUController) UpdateDataById() {
	objectId := ctx.GetString("objectId")

	update := ctx.GetString("update")
	var updateData models.UpdateMap
	err := json.Unmarshal([]byte(update), &updateData)
	if err != nil {
		ctx.ApiError(400, "update格式错误")
	}

	mp := ctx.ApiMongoProxy()
	object, err := mp.UpdateDataById(objectId, updateData)

	if err != nil {
		ctx.ApiFailData(1, "更新失败"+fmt.Sprintf("%s", err), object)
	}

	ctx.ApiSuccessData("更新成功", object)

}

// @Title 更新一份文档
// @Description 精确查找更新一个文档, {"ret": 200, "msg": "", "data": {"code": 0, "msg": "更新成功", "data": {"MatchedCount": 1, "ModifiedCount": 1, "UpsertedCount": 0, "UpsertedID": null}}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param    filter  		query	string  true		"条件更新文档的条件，json传递，如{"name":"hxh"}"
// @Param    update  		query	string  true		"条件更新文档的数据，json传递，更新的字段必须是filter中有的，如{"name":"hxh"}"
// @Success 0 {string} 更新成功
// @Failure 1 {string} 更新失败
// @router /UpdateOneData [post]
func (ctx *MongoUController) UpdateOneData() {

	filter := ctx.GetString("filter")
	update := ctx.GetString("update")
	var filterWhere models.FilterMap
	var updateData models.UpdateMap
	err1 := json.Unmarshal([]byte(filter), &filterWhere)
	if err1 != nil {
		ctx.ApiError(400, "filter格式错误")
	}

	err2 := json.Unmarshal([]byte(update), &updateData)
	if err2 != nil {
		ctx.ApiError(400, "update格式错误")
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.UpdateOneData(filterWhere, updateData)

	if err != nil {
		ctx.ApiFailData(1, "更新失败"+fmt.Sprintf("%s", err), result)
	}

	ctx.ApiSuccessData("更新成功", result)
}

// @Title 更新多份文档
// @Description 精确查找更新多份文档,{"ret": 200, "msg": "", "data": {"code": 0, "msg": "更新成功", "data": {"MatchedCount": 3, "ModifiedCount": 3, "UpsertedCount": 0, "UpsertedID": null}}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   filter  		query	string  true		"条件更新文档的条件,{"name":"hxh"}"
// @Param   update  		query	string  true		"条件更新文档的数据,更新的字段必须是filter中有的,{"name":"hxh"}"
// @Success 0 {string} 更新成功
// @Failure 1 {string} 更新失败
// @router /MultiUpdateData [post]
func (ctx *MongoUController) MultiUpdateData() {

	filter := ctx.GetString("filter")
	var filterWhere models.FilterMap
	err1 := json.Unmarshal([]byte(filter), &filterWhere)
	if err1 != nil {
		ctx.ApiError(400, "filter格式错误")
	}

	update := ctx.GetString("update")
	var updateData models.UpdateMap
	err2 := json.Unmarshal([]byte(update), &updateData)
	if err2 != nil {
		ctx.ApiError(400, "update格式错误")
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.MultiUpdateData(filterWhere, updateData)

	if err != nil {
		ctx.ApiFailData(1, "更新失败"+fmt.Sprintf("%s", err), result)
	}
	ctx.ApiSuccessData("更新成功", result)
}

// @Title 替换一份文档
// @Description 精确查找替换一份文档，这是替换整一份文档，不是去更新某个域，{"ret": 200, "msg": "", "data": {"code": 0, "msg": "替换成功", "data": {"MatchedCount": 1, "ModifiedCount": 1, "UpsertedCount": 0, "UpsertedID": null}}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param    filter  		query	string  true		"条件更新文档的条件,{"name":"hxh"}"
// @Param    update  		query	string  true		"条件更新文档的数据,{"name":"hxh"}"
// @Success 0 {string} 替换成功
// @Failure 1 {string} 替换失败
// @router /ReplaceOneData [post]
func (ctx *MongoUController) ReplaceOneData() {
	filter := ctx.GetString("filter")
	var filterWhere models.FilterMap
	err1 := json.Unmarshal([]byte(filter), &filterWhere)
	if err1 != nil {
		ctx.ApiError(400, "filter格式错误")
	}

	update := ctx.GetString("update")
	var updateData models.UpdateMap
	err2 := json.Unmarshal([]byte(update), &updateData)
	if err2 != nil {
		ctx.ApiError(400, "update格式错误")
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.ReplaceOneData(filterWhere, updateData)

	if err != nil {
		ctx.ApiFailData(1, "替换失败", result)
	}

	ctx.ApiSuccessData("替换成功", result)
}

// @Title 根据_id替换一份文档
// @Description 根据对象ID精确查找替换一份文档，这是替换整一份文档，不是去更新某个域，{"ret": 200, "msg": "", "data": {"code": 0, "msg": "替换成功", "data": {"MatchedCount": 1, "ModifiedCount": 1, "UpsertedCount": 0, "UpsertedID": null}}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   objectId		query	string   true		"对象ID"
// @Param   update  		query	string  true		"条件更新文档的数据,{"name":"hxh"}"
// @Success 0 {string} 替换成功
// @Failure 1 {string} 替换失败
// @router /ReplaceDataById [post]
func (ctx *MongoUController) ReplaceDataById() {
	objectId := ctx.GetString("objectId")
	update := ctx.GetString("update")
	var updateData models.UpdateMap
	err2 := json.Unmarshal([]byte(update), &updateData)
	if err2 != nil {
		ctx.ApiError(400, "update格式错误")
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.ReplaceDataById(objectId, updateData)

	if err != nil {
		ctx.ApiFailData(1, "替换失败"+fmt.Sprintf("%s", err), result)
	}

	ctx.ApiSuccessData("替换成功", result)
}
