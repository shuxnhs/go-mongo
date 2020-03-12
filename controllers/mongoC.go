package controllers

import (
	"encoding/json"
	"fmt"
	"go-mongo/models"
)

// mongo文档新增操作相关接口
type MongoCController struct {
	BaseController
}

// @Title Creating Objects
// @Description 为collection添加一条文档，{"ret": 200, "msg": "", "data": {"code": 0, "msg": "新增成功", "data": {"InsertedID": "5e6a6fbdce6ee32678ad1a95"}}}
// @Param	mongoKey		query 	string	true		"用户mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param	document		query 	string	true		"要插入的文档，JSON格式"
// @Success 0 {string} 添加成功，返回对象id,InsertedID：_id
// @Failure 1 {string} 添加失败
// @router /CreateOneDocument [post]
func (ctx *MongoCController) CreateOneDocument() {
	document := ctx.GetString("document")
	var ob models.Document
	err := json.Unmarshal([]byte(document), &ob)
	if err != nil {
		ctx.ApiError(400, "document的JSON格式错误")
	}

	mp := ctx.ApiMongoProxy()
	objectId, cErr := mp.CreateObject(ob)
	if cErr != nil {
		ctx.ApiFail(1, fmt.Sprintf("%s", err))
	}
	ctx.ApiSuccessData("新增成功", objectId)
}

// @Title 批量添加文档数据
// @Description 为collection批量添加文档，{"ret": 200, "msg": "", "data": {"code": 0, "msg": "添加成功", "data": {"InsertedIDs": ["5e6a71a42a44314bfb004fbe", "5e6a71a42a44314bfb004fbf"]}}}
// @Param	mongoKey		query 	string	true		"用户mongoKey"
// @Param   collection  	query	string  true		"集合名"
// @Param	documents		query 	string	true		"要批量插入的文档"
// @Success 0 {string} 添加成功，返回对象id的数组,InsertedID：_id
// @Failure 1 {string} 添加失败
// @router /MultiCreateDocuments [post]
func (ctx *MongoCController) MultiCreateDocuments() {
	documentString := ctx.GetString("documents")
	var documentList []interface{}
	err := json.Unmarshal([]byte(documentString), &documentList)
	if err != nil {
		ctx.ApiError(400, "documents格式错误")
	}
	mp := ctx.ApiMongoProxy()
	result, cErr := mp.MultiCreateData(documentList)
	if cErr != nil {
		ctx.ApiFailData(1, "批量创建失败"+fmt.Sprintf("%s", err), result)
	}
	ctx.ApiSuccessData("添加成功", result)
}
