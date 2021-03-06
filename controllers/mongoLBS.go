package controllers

import "fmt"

// mongoLBS操作相关接口->（附近的人）
type MongoLBSController struct {
	BaseController
}

// @Title 查询从近到远的地点
// @Description 查询最近到远的地点，注意集合必须先创建好索引2DSphere，返回的distance为坐标的距离，{"ret": 200, "msg": "", "data": {"code": 0, "msg": "查询成功", "data": [{"_id": "5e6f236a436ec50abd7457d4", "account": "1", "collectTime": 1480602671, "distance": 902576.4266556037, "location": {"coordinates": [107.840974298098, 33.2789316522934], "type": "Point"}, "logTime": 1480602675, "platform": "android"}]}}
// @Param	mongoKey		query 	string	true		"mongoKey"
// @Param   collection		query	string  true		"集合名"
// @Param   lon     		query	float   true		"查询的经度"
// @Param   lat     		query	float   true		"查询的纬度"
// @Param   maxDistance   	query	int     true		"查询多少距离内的，单位为米，默认1000000米"
// @Param   minDistance   	query	int     true		"查询最小多少距离外的，单位为米，默认1000米"
// @Param   num				query	int64   true		"获取条数"
// @router /GetNearAndDistance [get]
func (ctx *MongoLBSController) GetNearAndDistance() {
	lon, err := ctx.GetFloat("lon")
	if err != nil {
		ctx.ApiError(400, "lon参数错误")
	}
	lat, err := ctx.GetFloat("lat")
	if err != nil {
		ctx.ApiError(400, "lon参数错误")
	}
	maxDistance, err1 := ctx.GetInt64("maxDistance")
	if err1 != nil || maxDistance < 1 {
		maxDistance = 1000000
	}

	minDistance, err1 := ctx.GetInt64("minDistance")
	if err1 != nil || maxDistance < 1 {
		minDistance = 1000
	}

	num, err1 := ctx.GetInt64("num")
	if err1 != nil || num < 1 {
		ctx.ApiError(400, "num参数错误")
	}

	mp := ctx.ApiMongoProxy()
	result, err := mp.FindNearLBS(lon, lat, maxDistance, minDistance, num)

	if err != nil {
		ctx.ApiFailData(2, fmt.Sprintf("%s", err), result)
	}

	if len(result) == 0 {
		ctx.ApiFail(1, "集合文档为空")
	}
	ctx.ApiSuccessData("查询成功", result)
}

// 新增的文档：{"account" : "1", "platform" : "android", "location" : {"type" : "Point", "coordinates" : [108.840974298098, 34.2789316522934]}, "collectTime" : 1480602671, "logTime" : 1480602675}

// @Title 创建2DSphere索引接口
// @Description 建2DSphere索引,注意索引的key是Geojson类型{"type" : "Point", "coordinates" : [108.8098, 34.2789]}  {"ret": 200, "msg": "", "data": {"code": 0, "msg": "创建成功", "data": "1"}}
// @Param	mongoKey		query 	string	 true		"mongoKey"
// @Param   collection		query	string   true		"集合名"
// @Param   key     		query	string   true		"创建的索引字段"
// @Param   indexName		query	string   true		"索引名称"
// @router /Create2DSphereIndex [get]
func (ctx *MongoLBSController) Create2DSphereIndex() {
	key := ctx.GetString("key")
	indexName := ctx.GetString("indexName")

	mp := ctx.ApiMongoProxy()
	result, err := mp.Create2DSphereIndex(key, indexName)

	if err != nil {
		ctx.ApiFailData(2, fmt.Sprintf("%s", err), result)
	}

	ctx.ApiSuccessData("创建成功", result)
}
