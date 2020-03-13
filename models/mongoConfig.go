package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	utiLog "go-mongo/common/log"
	"go-mongo/common/rsa"
)

/**
 * 每个项目的mongodb的配置
 */

type Config struct {
	Id        int    `json:"id"`
	Mongo_key string `json:"mongo_key"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Dbname    string `json:"dbname"`
}

func AddMongoConfigWith(mk string, host string, port int, user string, pass string, dbname string) error {
	config := Config{
		Mongo_key: mk,
		Host:      host,
		Port:      port,
		User:      user,
		Password:  string(rsa.RSA_Encrypt([]byte(pass), "./common/public.pem")),
		Dbname:    dbname,
	}
	// 密码没有时，存空
	if pass == "" {
		config.Password = ""
	}
	return AddMongoConfig(config)
}

func AddMongoConfig(config Config) error {
	o := orm.NewOrm()
	_, err := o.Insert(&config)
	return err
}

func GetMongoConfig(Mk string) (Config, error) {
	o := orm.NewOrm()
	config := Config{}
	if gcache.IsExist(Mk) {
		if config, ok := gcache.Get(Mk).(Config); ok {
			utiLog.Log.Info("get cache config: ", config)
			return config, nil
		}
	}
	exist := o.QueryTable("Config").Filter("Mongo_key", Mk).Exist()
	if exist {
		_ = o.QueryTable("Config").Filter("Mongo_key", Mk).One(&config)
		config.Password = string(rsa.RSA_Decrypt([]byte(config.Password), "./common/private.pem"))
		if err := gcache.Put(Mk, config, 3600); err != nil {
			utiLog.Log.Error("cache the mk: ", Mk, "config err: ", err)
		}
		utiLog.Log.Info("cache the mk: ", Mk, " config: ", config)
		return config, nil
	}
	return config, errors.New("MongoConfig No Exist")
}

// 判断是否存在
func ExistMongoConfig(Mk string) bool {
	o := orm.NewOrm()
	exist := o.QueryTable("Config").Filter("Mongo_key", Mk).Exist()
	return exist
}

func DeleteMongoConfig(Mk string) (bool, error) {
	o := orm.NewOrm()
	exist := o.QueryTable("Config").Filter("Mongo_key", Mk).Exist()
	if exist {
		_, err := o.QueryTable("Config").Filter("Mongo_key", Mk).Delete()
		return true, err
	} else {
		return false, nil
	}
}

func UpdateMongoConfigWith(id int, host string, port int, user string, pass string, dbname string) error {
	o := orm.NewOrm()
	config := Config{Id: id}
	if o.Read(&config) == nil {
		config.Host = host
		config.Port = port
		config.User = user
		if pass != "" {
			config.Password = string(rsa.RSA_Encrypt([]byte(pass), "./common/public.pem"))
		}
		config.Dbname = dbname

		_, err := o.Update(&config)
		return err
	}
	return nil
}
