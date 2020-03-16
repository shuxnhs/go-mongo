package controllers

import "fmt"

// mongo全文搜索相关操作
type MongoFTController struct {
	BaseController
}

// @Title 全文搜索接口
// @Description 根据关键字全文搜索，注意集合必须先创建好索引，注意mongodb版本支持的语言,{"ret": 200, "msg": "", "data": {"code": 0, "msg": "查询成功", "data": [{"_id": "5e6e4617291c5d88fed24286", "age": 21, "name": "hxh"},{"_id": "5da368b34a0bab8c655a8142", "age": 1, "name": "hxh is a boy"}]}}(text=hxh)
// @Param	mongoKey		query 	string	 true		"mongoKey"
// @Param   collection		query	string   true		"集合名"
// @Param   text     		query	string   true		"搜索关键词"
// @Param   num				query	int64    false		"获取条数"
// @router /FullTextSearch [get]
func (ctx *MongoFTController) FullTextSearch() {
	text := ctx.GetString("text")

	num, err1 := ctx.GetInt64("num")
	if err1 != nil || num < 1 {
		ctx.ApiError(400, "num参数错误")
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.FullTextFind(text, num)

	if err != nil {
		ctx.ApiFailData(2, fmt.Sprintf("%s", err), result)
	}

	if len(result) == 0 {
		ctx.ApiFail(1, "集合文档为空")
	}
	ctx.ApiSuccessData("查询成功", result)
}

// @Title 创建全文索引接口
// @Description 创建全文索引,{"ret": 200, "msg": "", "data": {"code": 0, "msg": "创建成功", "data": "1"}}
// @Param	mongoKey		query 	string	 true		"mongoKey"
// @Param   collection		query	string   true		"集合名"
// @Param   key     		query	string   true		"创建的索引字段"
// @Param   indexName		query	string   true		"索引名称"
// @router /CreateFullText [get]
func (ctx *MongoFTController) CreateFullText() {
	key := ctx.GetString("key")
	indexName := ctx.GetString("indexName")

	mp := ctx.ApiMongoProxy()
	result, err := mp.CreateFullTextIndex(key, indexName)

	if err != nil {
		ctx.ApiFailData(2, fmt.Sprintf("%s", err), result)
	}

	ctx.ApiSuccessData("创建成功", result)
}
