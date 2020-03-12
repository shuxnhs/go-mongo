package models

import (
	"errors"
	utiLog "go-mongo/common/log"
	"sync"
)

type mongoConnectionManager struct {
	conn     map[string]MongoConnection
	connLock sync.RWMutex // 保护链接的读写锁
}

// 创建方法
func NewMongoProxyManager() *mongoConnectionManager {
	return &mongoConnectionManager{
		conn: make(map[string]MongoConnection),
	}
}

// 添加链接
func (m *mongoConnectionManager) AddConn(mongoConnection *MongoConnection) {
	// 加写锁，保护共享资源
	m.connLock.Lock()
	defer m.connLock.Unlock()
	m.conn[mongoConnection.MK] = *mongoConnection
	utiLog.Log.Info("mongo-key: ", mongoConnection.MK, " add to mongoProxyManager")
}

// 删除链接
func (m *mongoConnectionManager) DelConn(mongoConnection *MongoConnection) {
	// 加写锁，保护共享资源
	m.connLock.Lock()
	defer m.connLock.Unlock()
	_ = mongoConnection.CloseMongo()
	delete(m.conn, mongoConnection.MK)
	utiLog.Log.Info("mongo-key: ", mongoConnection.MK, " delete from mongoProxyManager")
}

// 根据MK获取链接
func (m *mongoConnectionManager) GetConn(mk string) (*MongoConnection, error) {
	// 加读锁，保护共享资源
	m.connLock.RLock()
	defer m.connLock.RUnlock()
	if conn, ok := m.conn[mk]; ok {
		if err := conn.CheckPing(); err != nil {
			m.DelConn(&conn)
			return nil, errors.New("conn timeout")
		}
		return &conn, nil
	}
	return nil, errors.New("Conn not Exist ")
}

// 获取所有链接的总数
func (m *mongoConnectionManager) GetConnNum() int {
	return len(m.conn)
}

// 删除所有链接
func (m *mongoConnectionManager) DelAllConn() {
	// 加写锁，保护共享资源
	m.connLock.Lock()
	defer m.connLock.Unlock()
	for mk, mongoConnection := range m.conn {
		// 停止conn
		_ = mongoConnection.CloseMongo()
		// 删除conn
		delete(m.conn, mk)
	}
	utiLog.Log.Info("all connection was clear")
}
