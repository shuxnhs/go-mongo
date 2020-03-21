package controllers

import (
	"testing"
)

func BenchmarkGoMongoPing(b *testing.B) {
	mongo := NewGoMongoSDK()
	mongo.SetMongoHost("http://127.0.0.1:8081").SetMongoKey("1DD75A62EB5E561F0F10A9A51270E5A6")
	for i := 0; i < 8000; i++ {
		res, err := mongo.Request("mongodb", "CheckMongoConnect", map[string]string{}, "GET")
		if err != nil {
			b.Logf("err : %s", err)
		}
		b.Log(res.Data)
	}
}
