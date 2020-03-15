package models

import (
	"context"
	"fmt"
	utiLog "go-mongo/common/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"time"
)

type MongoConnection struct {
	MK         string
	Database   string
	Collection string
	Mongo      *mongo.Client
}

type Document map[string]interface{}      // 文档
type FilterMap map[string]interface{}     // 过滤条件
type UpdateMap map[string]interface{}     // 更新条件
type ProjectionMap map[string]interface{} // 返回字段

// 实例化
func NewMongoConnection(mp *MongoProxy, Mk string, Collection string) *MongoConnection {
	config, err := GetMongoConfig(Mk)
	if err != nil {
		utiLog.Log.Error("get mongo config fail, err: ", err)
		return nil
	}
	whoAmI := ""
	if config.User != "" {
		whoAmI = config.User + ":" + url.QueryEscape(config.Password) + "@"
	}
	connectUrl := "mongodb://" + whoAmI + config.Host + ":" + fmt.Sprintf("%d", config.Port)

	// Set client options
	context.WithTimeout(context.Background(), 10*time.Second)

	clientOptions.ApplyURI(connectUrl).SetAppName(Mk)
	// Connect to MongoDB
	Mongo, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		utiLog.Log.Error("mongo client connect fail, err: ", err)
		return nil
	}
	// Check the connection
	err = Mongo.Ping(context.TODO(), nil)
	if err != nil {
		utiLog.Log.Error("mongo client ping fail, err: ", err)
		return nil
	}
	conn := &MongoConnection{
		MK:         Mk,
		Database:   config.Dbname,
		Collection: Collection,
		Mongo:      Mongo,
	}
	mp.GetMongoProxyManager().AddConn(conn)
	return conn
}

// 辅助的方法，切换数据库
func (mc *MongoConnection) SwitchDatabase(Database string) *MongoConnection {
	mc.Database = Database
	return mc
}

// 切换集合
func (mc *MongoConnection) SwitchCollection(Collection string) *MongoConnection {
	mc.Collection = Collection
	return mc
}

// 获取指定数据库和集合的句柄
func (mc *MongoConnection) CurCollection() *mongo.Collection {
	return mc.Mongo.Database(mc.Database).Collection(mc.Collection)
}

// 检查是否连接
func (mc *MongoConnection) CheckPing() error {
	err := mc.Mongo.Ping(context.TODO(), nil)
	if err != nil {
		utiLog.Log.Error("mongo client ping fail, err: ", err)
		return err
	}
	return nil
}

// 关闭连接
func (mc *MongoConnection) CloseMongo() error {
	if mc.Mongo == nil {
		return nil
	}
	err := mc.Mongo.Disconnect(context.TODO())
	if err != nil {
		utiLog.Log.Error("mongo client close fail, err: ", err)
		return err
	}
	utiLog.Log.Info("mongo client close success")
	return nil
}

/**-----------------------------获取文档操作---------------------------------**/

/**
 * @func: 获取文档数量
 */
func (mc *MongoConnection) CountData(filter FilterMap) (int64, error) {
	return mc.CurCollection().CountDocuments(mc.getContext(), filter)
}

/**
 * @func: 根据id获取文档
 */
func (mc *MongoConnection) RetrieveObject(objectId string) (Document, error) {
	var document Document
	id, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return document, err
	}
	err = mc.CurCollection().FindOne(mc.getContext(), bson.M{"_id": id}).Decode(&document)
	return document, err
}

/**
 * @func：随机n份文档
 */
func (mc *MongoConnection) GetDataList(num int64) ([]*Document, error) {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(num)

	// Here's an array in which you can store the decoded documents
	var results []*Document

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := mc.CurCollection().Find(mc.getContext(), bson.D{{}}, findOptions)

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		var elem Document
		_ = cur.Decode(&elem)
		results = append(results, &elem)
	}
	// Close the cursor once finished
	_ = cur.Close(context.TODO())
	return results, err
}

/**
 * @func：自由获取一条文档
 */
func (mc *MongoConnection) FreeFindOne(filter FilterMap) (Document, error) {
	filterData, _ := bson.Marshal(filter)
	var document Document
	err := mc.CurCollection().FindOne(mc.getContext(), filterData).Decode(&document)
	return document, err
}

/**
 * @func：自由获取n份文档
 */
func (mc *MongoConnection) FreeGetDataList(projectionOpt ProjectionMap, filter FilterMap, num int64) ([]*Document, error) {
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetProjection(projectionOpt)
	findOptions.SetLimit(num)
	// findOptions.SetSkip(1)   // 相当于offset
	// findOptions.SetSort(bson.M{"createtime": -1}) // 1升序 -1降序

	// Here's an array in which you can store the decoded documents
	var results []*Document

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := mc.CurCollection().Find(mc.getContext(), filter, findOptions)

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		var elem Document
		_ = cur.Decode(&elem)
		results = append(results, &elem)
	}
	// Close the cursor once finished
	_ = cur.Close(context.TODO())
	return results, err
}

/**
 * @func：获取大于，等于，小于等条件的n条数据
 */
func (mc *MongoConnection) CondOperateFind(cond string, key string, value int64, num int64) ([]*Document, error) {
	operate := mc.condOperateChange(cond)
	findOptions := options.Find().SetLimit(num)
	var results []*Document
	cur, err := mc.CurCollection().Find(mc.getContext(), bson.M{key: bson.M{operate: value}}, findOptions)
	defer cur.Close(mc.getContext())
	for cur.Next(mc.getContext()) {
		var elem Document
		_ = cur.Decode(&elem)
		results = append(results, &elem)
	}
	return results, err
}

/**
 * @func: 条件操作符算有多少条数据
 */
func (mc *MongoConnection) CondOperateCount(cond string, key string, value int64) (int64, error) {
	operate := mc.condOperateChange(cond)
	return mc.CurCollection().CountDocuments(mc.getContext(), bson.M{key: bson.M{operate: value}})
}

/**
 * @func：获取字段类型的条件的n条数据
 * Double =>    1
 * String =>	2
 * Object => 	3
 * Array  =>	4
 * Binary data	 =>   5
 * Undefined	 =>   6	 已废弃。
 * Object id	 =>   7
 * Boolean	     =>   8
 * Date	  =>  9
 * Null	  =>  10
 * Regular Expression   =>	11
 * JavaScript	        =>  13
 * Symbol	            =>  14
 * JavaScript (with scope)	 =>   15
 * 32-bit integer	    =>   16
 * Timestamp			=>   17
 * 64-bit integer	    =>   18
 * Min key	  255	Query with -1.
 * Max key	  127
 */
func (mc *MongoConnection) TypeOperateFind(typeKey int8, key string, num int64) ([]*Document, error) {
	findOptions := options.Find().SetLimit(num)
	var results []*Document
	cur, err := mc.CurCollection().Find(mc.getContext(), bson.M{key: bson.M{"$type": typeKey}}, findOptions)
	defer cur.Close(mc.getContext())
	for cur.Next(mc.getContext()) {
		var elem Document
		_ = cur.Decode(&elem)
		results = append(results, &elem)
	}
	return results, err
}

/**---------------------------更新文档操作-------------------------------**/

/**
 *@func: 根据_id更新数据
 */
func (mc *MongoConnection) UpdateDataById(objectId string, update UpdateMap) (*mongo.UpdateResult, error) {
	id, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return &mongo.UpdateResult{
			MatchedCount:  0,
			ModifiedCount: 0,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}, err
	}
	return mc.CurCollection().UpdateOne(mc.getContext(), bson.M{"_id": id}, bson.D{{"$set", update}})
}

/**
 *@func: 自由更新一条数据
 */
func (mc *MongoConnection) UpdateOneData(filter FilterMap, update UpdateMap) (*mongo.UpdateResult, error) {
	filterData, err := bson.Marshal(filter)
	if err != nil {
		return &mongo.UpdateResult{
			MatchedCount:  0,
			ModifiedCount: 0,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}, err
	}
	return mc.CurCollection().UpdateOne(mc.getContext(), filterData, bson.D{{"$set", update}})
}

/**
 * @func：批量更新数据
 */
func (mc *MongoConnection) MultiUpdateData(filter FilterMap, update UpdateMap) (*mongo.UpdateResult, error) {
	filterData, err := bson.Marshal(filter)
	if err != nil {
		return &mongo.UpdateResult{
			MatchedCount:  0,
			ModifiedCount: 0,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}, err
	}
	return mc.CurCollection().UpdateMany(mc.getContext(), filterData, bson.D{{"$set", update}})
}

/**
 * @func: 替换数据
 */
func (mc *MongoConnection) ReplaceOneData(filter FilterMap, update UpdateMap) (*mongo.UpdateResult, error) {
	filterData, err := bson.Marshal(filter)
	if err != nil {
		return &mongo.UpdateResult{
			MatchedCount:  0,
			ModifiedCount: 0,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}, err
	}
	return mc.CurCollection().ReplaceOne(mc.getContext(), filterData, update)
}

/**
 *@func: 根据_id替换数据
 */
func (mc *MongoConnection) ReplaceDataById(objectId string, update UpdateMap) (*mongo.UpdateResult, error) {
	id, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return &mongo.UpdateResult{
			MatchedCount:  0,
			ModifiedCount: 0,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}, err
	}
	return mc.CurCollection().ReplaceOne(mc.getContext(), bson.M{"_id": id}, update)
}

/**---------------------------删除文档操作-------------------------------**/

/**
 * @func: 根据_id删除文档
 */
func (mc *MongoConnection) DeleteDataById(objectId string) (*mongo.DeleteResult, error) {
	id, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		return &mongo.DeleteResult{DeletedCount: int64(0)}, err
	}
	return mc.CurCollection().DeleteOne(mc.getContext(), bson.M{"_id": id})
}

/**
 * @func: 删除文档
 */
func (mc *MongoConnection) DeleteOneData(filter FilterMap) (*mongo.DeleteResult, error) {
	filterData, err := bson.Marshal(filter)
	if err != nil {
		return &mongo.DeleteResult{DeletedCount: int64(0)}, err
	}
	return mc.CurCollection().DeleteOne(mc.getContext(), filterData)
}

/**
 * @func: 批量删除文档
 */
func (mc *MongoConnection) MultiDeleteData(filter FilterMap) (*mongo.DeleteResult, error) {
	filterData, err := bson.Marshal(filter)
	if err != nil {
		return &mongo.DeleteResult{DeletedCount: int64(0)}, err
	}
	return mc.CurCollection().DeleteMany(mc.getContext(), filterData)
}

/**---------------------------增加文档操作-------------------------------**/

/**
 * @func: 插入新的文档
 */
func (mc *MongoConnection) CreateObject(document Document) (*mongo.InsertOneResult, error) {
	return mc.CurCollection().InsertOne(mc.getContext(), document)
}

/**
 * @func: 批量新增文档
 */
func (mc *MongoConnection) MultiCreateData(documentList []interface{}) (*mongo.InsertManyResult, error) {
	return mc.CurCollection().InsertMany(mc.getContext(), documentList)
}

/****---------------------------------辅助方法--------------------------------------****/

/**
 * @func: 条件运算符转换
 */
func (mc *MongoConnection) condOperateChange(cond string) string {
	var operate string
	switch cond {
	case ">":
		operate = "$gt"
	case "<":
		operate = "$lt"
	case ">=":
		operate = "$gte"
	case "<=":
		operate = "$lte"
	case "!=":
		operate = "$ne"
	default:
		operate = "$eq"
	}
	return operate
}

/**
 * @func: context统一控制
 */
func (mc *MongoConnection) getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	return ctx
}
