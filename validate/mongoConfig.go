package validate

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"reflect"
)

//tags中 alias 表示验证不通过时候的提示名称，valid表示验证的格式 分号分割不同的验证内容
type ValidateMongoConfig struct {
	MongoKey string `alias:"mongo_key"  valid:"Required;Length(32)"`
	Host     string `alias:"mongo的host"  valid:"Required;IP"`
}

// 验证添加配置的参数
func (a *ValidateMongoConfig) ValidAddMongo() interface{} {
	valid := validation.Validation{}
	b, _ := valid.Valid(a)
	if !b {
		//表示获取验证的结构体
		st := reflect.TypeOf(ValidateMongoConfig{})
		for _, err := range valid.Errors {
			//获取验证的字段名和提示信息的别名
			filed, _ := st.FieldByName(err.Field)
			var alias = filed.Tag.Get("alias")
			//返回验证的错误信息
			errMsg := fmt.Sprintf("%s", errors.New(alias+err.Message))
			return errMsg
		}
	}
	return nil
}

//验证获取和删除的参数
type ValidateMk struct {
	MongoKey string `alias:"mongo_key"  valid:"Required;Length(32)"`
}

func (a *ValidateMk) ValidDelOrGetMongo() interface{} {
	valid := validation.Validation{}
	b, _ := valid.Valid(a)
	if !b {
		//表示获取验证的结构体
		st := reflect.TypeOf(ValidateMk{})
		for _, err := range valid.Errors {
			//获取验证的字段名和提示信息的别名
			filed, _ := st.FieldByName(err.Field)
			var alias = filed.Tag.Get("alias")
			//返回验证的错误信息
			errMsg := fmt.Sprintf("%s", errors.New(alias+err.Message))
			return errMsg
		}
	}
	return nil
}
