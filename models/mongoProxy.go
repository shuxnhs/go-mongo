package models

import utiLog "go-mongo/common/log"

type MongoProxy struct {
	MongoConnectionManager *mongoConnectionManager
}

func NewMongoProxy() *MongoProxy {
	return &MongoProxy{
		MongoConnectionManager: NewMongoProxyManager(),
	}
}

func (m *MongoProxy) Start(Mk string, Collection string) *MongoConnection {
	utiLog.Log.Info("start new mongo connect")
	return NewMongoConnection(m, Mk, Collection)
}

// 获取当前连接管理器
func (m *MongoProxy) GetMongoProxyManager() *mongoConnectionManager {
	return m.MongoConnectionManager
}
