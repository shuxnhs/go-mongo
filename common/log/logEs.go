package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OwnLocal/goes"
	"github.com/astaxie/beego/logs"
	"net"
	"net/url"
	"time"
)

// NewES return a LoggerInterface
func NewES() logs.Logger {
	cw := &esLogger{
		Level: logs.LevelDebug,
	}
	return cw
}

type esLogger struct {
	Client *goes.Connection
	DSN    string `json:"dsn"`
	Level  int    `json:"level"`
}

// {"dsn":"http://localhost:9200/","level":1}
func (el *esLogger) Init(jsonconfig string) error {
	err := json.Unmarshal([]byte(jsonconfig), el)
	if err != nil {
		return err
	}
	if el.DSN == "" {
		return errors.New("empty dsn")
	} else if u, err := url.Parse(el.DSN); err != nil {
		return err
	} else if u.Path == "" {
		return errors.New("missing prefix")
	} else if host, port, err := net.SplitHostPort(u.Host); err != nil {
		return err
	} else {
		conn := goes.NewConnection(host, port)
		el.Client = conn
	}
	return nil
}

// WriteMsg will write the msg and level into es
func (el *esLogger) WriteMsg(when time.Time, msg string, level int) error {

	if level > el.Level {
		return nil
	}

	vals := make(map[string]interface{})
	vals["@timestamp"] = when.Format(time.RFC3339)
	vals["@msg"] = msg
	d := goes.Document{
		Index:  fmt.Sprintf("%v.%04d.%02d.%02d", "go-mongo-log", when.Year(), when.Month(), when.Day()),
		Type:   "logs",
		Fields: vals,
	}
	_, err := el.Client.Index(d, nil)
	return err
}

// Destroy is a empty method
func (el *esLogger) Destroy() {

}

// Flush is a empty method
func (el *esLogger) Flush() {

}
