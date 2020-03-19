package tests

import (
	"context"
	"fmt"
	"go-mongo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"testing"
	"time"
)

type fields struct {
	MK         string
	Database   string
	Collection string
	Mongo      *mongo.Client
}

// 只返回err的用例
type testSample struct {
	name    string
	fields  fields
	wantErr bool
}

// 返回err和res的用例
type testSampleWithInsertValue struct {
	name    string
	fields  fields
	args    interface{}
	want    interface{}
	wantErr bool
}

var Mk string
var Database string
var Collection string
var MongoClient *mongo.Client
var ObjectId string

// 测试初始化
func init() {
	Mk = "1DD75A62EB5E561F0F10A9A51270E5A6"
	Database = "goTest"
	Collection = "test-" + fmt.Sprintf("%s", time.Now().Format("2006-01-02")) // 每次运行test都新建一个Collection，暂时按天新建
	connectUrl := "mongodb://127.0.0.1:27017"
	// Connect to MongoDB
	MongoClient, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI(connectUrl))
}

func TestMongoConnection_CheckPing(t *testing.T) {
	tests := []testSample{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			t.Log(mc.MK)
			if err := mc.CheckPing(); (err != nil) != tt.wantErr {
				t.Errorf("CheckPing() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongoConnection_CreateObject(t *testing.T) {
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			map[string]interface{}{
				"name": "hxh",
				"age":  1,
			},
			&mongo.InsertOneResult{InsertedID: nil},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			got, err := mc.CreateObject(tt.args.(map[string]interface{}))
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateObject() got = %v, no want %v", got, tt.want)
			}
			t.Logf("CreateObject() got = %v", got)
			ObjectId = got.InsertedID.(primitive.ObjectID).Hex()
		})
	}
}

func TestMongoConnection_MultiCreateData(t *testing.T) {
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			[]interface{}{
				map[string]interface{}{
					"name": "hxh",
					"age":  2,
				},
				map[string]interface{}{
					"name": "hxh",
					"age":  3,
				},
			},
			&mongo.InsertManyResult{InsertedIDs: []interface{}{}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			got, err := mc.MultiCreateData(tt.args.([]interface{}))
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiCreateData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiCreateData() got = %v, want %v", got, tt.want)
			}
			t.Logf("CreateObject() got = %v", got)
		})
	}
}

func TestMongoConnection_CondOperateCount(t *testing.T) {
	type Args struct {
		cond  string
		key   string
		value int64
	}
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			Args{
				cond:  ">",
				key:   "age",
				value: 2,
			},
			int64(0),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			args := tt.args.(Args)
			got, err := mc.CondOperateCount(args.cond, args.key, args.value)
			t.Log(got, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("CondOperateCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= tt.want.(int64) {
				t.Errorf("CondOperateCount() got = %v, noWant %v", got, tt.want)
			}
			t.Logf("CondOperateCount() got = %v", got)
		})
	}
}

func TestMongoConnection_CondOperateFind(t *testing.T) {
	type args struct {
		cond  string
		key   string
		value int64
		num   int64
	}
	Args := args{
		cond:  "=",
		key:   "age",
		value: int64(1),
		num:   1,
	}
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			Args,
			[]*models.Document{}, // 空的
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			Args := tt.args.(args)
			got, err := mc.CondOperateFind(Args.cond, Args.key, Args.value, Args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("CondOperateFind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("CondOperateFind() got = %v, want %v", got, tt.want)
			}
			if len(got) <= 0 {
				t.Errorf("CondOperateFind() got nil %v", got)
			}
			t.Logf("CondOperateFind() got = %v", got)
		})
	}
}

func TestMongoConnection_CountData(t *testing.T) {
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			map[string]interface{}{
				"name": "hxh",
			},
			int64(0),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			got, err := mc.CountData(tt.args.(map[string]interface{}))
			if (err != nil) != tt.wantErr {
				t.Errorf("CountData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got <= tt.want.(int64) {
				t.Errorf("CountData() got = %v, want %v", got, tt.want)
			}
			t.Logf("CondOperateCount() got = %v", got)
		})
	}
}

func TestMongoConnection_TypeOperateFind(t *testing.T) {
	type args struct {
		typeKey int8
		key     string
		num     int64
	}
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			args{
				typeKey: 2,
				key:     "name",
				num:     1,
			},
			[]*models.Document{}, // 空的
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			args := tt.args.(args)
			got, err := mc.TypeOperateFind(args.typeKey, args.key, args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeOperateFind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeOperateFind() got = %v, no want %v", got, tt.want)
			}
			t.Logf("TypeOperateFind() got = %v", got)
		})
	}
}

func TestMongoConnection_FreeFindOne(t *testing.T) {
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			map[string]interface{}{
				"name": "hxh",
			},
			models.Document{}, // 空的
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			got, err := mc.FreeFindOne(tt.args.(map[string]interface{}))
			if (err != nil) != tt.wantErr {
				t.Errorf("FreeFindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreeFindOne() got = %v, want %v", got, tt.want)
			}
			if len(got) <= 0 {
				t.Errorf("FreeFindOne() got no document %v", got)
			}
			t.Logf("FreeFindOne() got = %v", got)
		})
	}
}

func TestMongoConnection_FreeGetDataList(t *testing.T) {
	type args struct {
		projectionOpt models.ProjectionMap
		filter        models.FilterMap
		num           int64
	}
	Args := args{
		projectionOpt: nil,
		filter: map[string]interface{}{
			"name": "hxh",
		},
		num: int64(1),
	}

	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			Args,
			[]*models.Document{}, // 空的
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			Args = tt.args.(args)
			got, err := mc.FreeGetDataList(Args.projectionOpt, Args.filter, Args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("FreeGetDataList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreeGetDataList() got = %v, noWant %v", got, tt.want)
			}
			if len(got) <= 0 {
				t.Errorf("FreeGetDataList() got no document %v", got)
			}
			t.Logf("FreeGetDataList() got = %v", got)
		})
	}
}

func TestMongoConnection_GetDataList(t *testing.T) {
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			int64(1),
			[]*models.Document{}, // 空的
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			got, err := mc.GetDataList(tt.args.(int64))
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataList() got = %v, no want %v", got, tt.want)
			}
			if len(got) <= 0 {
				t.Errorf("GetDataList() got nil %v", got)
			}
			t.Logf("GetDataList() got = %v", got)
		})
	}
}

// 依赖新增文档的测试提供的ObjectId， 单独跑会不通过
func TestMongoConnection_RetrieveObject(t *testing.T) {
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			ObjectId,
			[]*models.Document{}, // 空的
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			t.Log(ObjectId)
			got, err := mc.RetrieveObject(tt.args.(string))
			if (err != nil) != tt.wantErr {
				t.Errorf("RetrieveObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("RetrieveObject() got = %v, no want %v", got, tt.want)
			}
			t.Logf("RetrieveObject() got = %v", got)
		})
	}
}

func TestMongoConnection_DeleteDataById(t *testing.T) {
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			ObjectId,
			&mongo.DeleteResult{DeletedCount: int64(1)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			got, err := mc.DeleteDataById(tt.args.(string))
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteDataById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteDataById() got = %v, want %v", got, tt.want)
			}
			t.Logf("DeleteDataById() got = %v", got)
		})
	}
}

func TestMongoConnection_DeleteOneData(t *testing.T) {
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			map[string]interface{}{
				"name": "hxh",
			},
			&mongo.DeleteResult{DeletedCount: int64(1)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			got, err := mc.DeleteOneData(tt.args.(map[string]interface{}))
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteOneData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteOneData() got = %v, want %v", got, tt.want)
			}
			t.Logf("DeleteOneData() got = %v", got)
		})
	}
}

func TestMongoConnection_MultiDeleteData(t *testing.T) {
	tests := []testSampleWithInsertValue{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			map[string]interface{}{
				"age": 2,
			},
			&mongo.DeleteResult{DeletedCount: int64(0)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			got, err := mc.MultiDeleteData(tt.args.(map[string]interface{}))
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiDeleteData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiDeleteData() got = %v, no want %v", got, tt.want)
			}
			if got.DeletedCount <= int64(0) {
				t.Errorf("MultiDeleteData() no delete %v", got)
			}
			t.Logf("MultiDeleteData() got = %v, delete num =%v", got, got.DeletedCount)
		})
	}
}

//func TestMongoConnection_Create2DSphereIndex(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		key       string
//		indexName string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.Create2DSphereIndex(tt.args.key, tt.args.indexName)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Create2DSphereIndex() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("Create2DSphereIndex() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_CreateFullTextIndex(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		key       string
//		indexName string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.CreateFullTextIndex(tt.args.key, tt.args.indexName)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CreateFullTextIndex() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("CreateFullTextIndex() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_CreateIndex(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		name         string
//		isBackground bool
//		isUnique     bool
//		weight       int
//		isSetSparse  bool
//		keys         []string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.CreateIndex(tt.args.name, tt.args.isBackground, tt.args.isUnique, tt.args.weight, tt.args.isSetSparse, tt.args.keys...)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CreateIndex() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("CreateIndex() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//

//func TestMongoConnection_CurCollection(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   *mongo.Collection
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			if got := mc.CurCollection(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CurCollection() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_FindNearLBS(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		lon         float64
//		lat         float64
//		maxDistance int64
//		minDistance int64
//		num         int64
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []*Document
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.FindNearLBS(tt.args.lon, tt.args.lat, tt.args.maxDistance, tt.args.minDistance, tt.args.num)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("FindNearLBS() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("FindNearLBS() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_FullTextFind(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		text string
//		num  int64
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []*Document
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.FullTextFind(tt.args.text, tt.args.num)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("FullTextFind() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("FullTextFind() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_MultiUpdateData(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		filter FilterMap
//		update UpdateMap
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *mongo.UpdateResult
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.MultiUpdateData(tt.args.filter, tt.args.update)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("MultiUpdateData() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("MultiUpdateData() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_ReplaceDataById(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		objectId string
//		update   UpdateMap
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *mongo.UpdateResult
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.ReplaceDataById(tt.args.objectId, tt.args.update)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ReplaceDataById() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ReplaceDataById() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_ReplaceOneData(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		filter FilterMap
//		update UpdateMap
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *mongo.UpdateResult
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.ReplaceOneData(tt.args.filter, tt.args.update)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ReplaceOneData() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ReplaceOneData() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_SwitchCollection(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		Collection string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *MongoConnection
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			if got := mc.SwitchCollection(tt.args.Collection); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("SwitchCollection() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_SwitchDatabase(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		Database string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *MongoConnection
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			if got := mc.SwitchDatabase(tt.args.Database); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("SwitchDatabase() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_UpdateDataById(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		objectId string
//		update   UpdateMap
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *mongo.UpdateResult
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.UpdateDataById(tt.args.objectId, tt.args.update)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("UpdateDataById() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("UpdateDataById() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_UpdateOneData(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		filter FilterMap
//		update UpdateMap
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *mongo.UpdateResult
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			got, err := mc.UpdateOneData(tt.args.filter, tt.args.update)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("UpdateOneData() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("UpdateOneData() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_condOperateChange(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	type args struct {
//		cond string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			if got := mc.condOperateChange(tt.args.cond); got != tt.want {
//				t.Errorf("condOperateChange() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestMongoConnection_getContext(t *testing.T) {
//	type fields struct {
//		MK         string
//		Database   string
//		Collection string
//		Mongo      *mongo.Client
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantCtx context.Context
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mc := &MongoConnection{
//				MK:         tt.fields.MK,
//				Database:   tt.fields.Database,
//				Collection: tt.fields.Collection,
//				Mongo:      tt.fields.Mongo,
//			}
//			if gotCtx := mc.getContext(); !reflect.DeepEqual(gotCtx, tt.wantCtx) {
//				t.Errorf("getContext() = %v, want %v", gotCtx, tt.wantCtx)
//			}
//		})
//	}
//}
//

//func TestNewMongoConnection(t *testing.T) {
//	type args struct {
//		mp         *models.MongoProxy
//		Mk         string
//		Collection string
//	}
//	tests := []struct {
//		name string
//		args args
//		want *models.MongoConnection
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewMongoConnection(tt.args.mp, tt.args.Mk, tt.args.Collection); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewMongoConnection() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestMongoConnection_CloseMongo(t *testing.T) {
	tests := []testSample{
		{
			"sample1",
			fields{
				Mk,
				Database,
				Collection,
				MongoClient,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &models.MongoConnection{
				MK:         tt.fields.MK,
				Database:   tt.fields.Database,
				Collection: tt.fields.Collection,
				Mongo:      tt.fields.Mongo,
			}
			if err := mc.CloseMongo(); (err != nil) != tt.wantErr {
				t.Errorf("CloseMongo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
