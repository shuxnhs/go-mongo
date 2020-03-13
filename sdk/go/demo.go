package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	mongo := NewGoMongoSDK()
	mongo.SetMongoHost("http://127.0.0.1:8081").SetMongoKey("")

	// get请求
	data := make(map[string]string)
	data["collection"] = "test1"
	data["filter"] = "{}"
	res, err := mongo.Request("mongoR", "CountData", data, "GET")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	// {200  map[code:0 data:3 msg:获取成功]}

	// post请求
	document := make(map[string]string)
	document["name"] = "go-mongo"
	byte, _ := json.Marshal(document)
	data = make(map[string]string)
	data["collection"] = "test1"
	data["document"] = string(byte)
	res, err = mongo.Request("mongoC", "CreateOneDocument", data, "POST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	// {200  map[code:0 data:map[InsertedID:5e6bb4b9b811645b654eba72] msg:新增成功]}
}
