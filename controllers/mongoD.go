package controllers

import (
	"encoding/json"
	"fmt"
	"go-mongo/models"
)

// mongo文档删除操作相关接口
type MongoDController struct {
	BaseController
}

// @Title 删除一份文档
// @Description 精确查找删除一份文档，{"ret": 200, "msg": "", "data": {"code": 0, "msg": "删除成功", "data": {"DeletedCount": 1}}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   filter  	    query	string  true		"删除文档的条件{"name":"hxh"}"
// @Success 0 {string} 清空成功
// @Failure 1 {string} 清空失败错误信息
// @router /DeleteOneData [post]
func (ctx *MongoDController) DeleteOneData() {
	filter := ctx.GetString("filter")
	var filterWhere models.FilterMap
	err1 := json.Unmarshal([]byte(filter), &filterWhere)
	if err1 != nil {
		ctx.ApiError(400, "filter格式错误")
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.DeleteOneData(filterWhere)
	if err != nil {
		ctx.ApiFailData(1, "删除失败"+fmt.Sprintf("%s", err), result)
	}
	ctx.ApiSuccessData("删除成功", result)
}

// @Title 删除多份文档
// @Description 精确查找删除多份文档,{"ret": 200, "msg": "", "data": {"code": 0, "msg": "删除成功", "data": {"DeletedCount": 2}}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   filter  	    query	string  true		"删除文档的条件{"name":"hxh"}"
// @Success 0 {string} 清空成功
// @Failure 1 {string} 清空失败错误信息
// @router /MultiDeleteData [post]
func (ctx *MongoDController) MultiDeleteData() {

	filter := ctx.GetString("filter")
	var filterWhere models.FilterMap
	err1 := json.Unmarshal([]byte(filter), &filterWhere)
	if err1 != nil {
		ctx.ApiError(400, "filter格式错误")
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.MultiDeleteData(filterWhere)

	if err != nil {
		ctx.ApiFailData(1, "删除失败"+fmt.Sprintf("%s", err), result)
	}

	ctx.ApiSuccessData("删除成功", result)
}

// @Title 根据_id删除一份文档
// @Description 根据对象精确查找删除一份文档，{"ret": 200, "msg": "", "data": {"code": 0, "msg": "删除成功", "data": {"DeletedCount": 1}}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   objectId		query	string   true		"对象ID"
// @Success 0 {string} 清空成功
// @Failure 1 {string} 清空失败错误信息
// @router /DeleteDataById [get]
func (ctx *MongoDController) DeleteDataById() {
	objectId := ctx.GetString("objectId")
	mp := ctx.ApiMongoProxy()
	result, err := mp.DeleteDataById(objectId)
	if err != nil {
		ctx.ApiFailData(1, "删除失败"+fmt.Sprintf("%s", err), result)
	}
	ctx.ApiSuccessData("删除成功", result)
}
