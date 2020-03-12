package validate

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"reflect"
)

//验证获取和删除的参数
type ValidateMongoKey struct {
	Mongo_key string `alias:"mongo_key" valid:"Required;Length(32)"`
}

func (v *ValidateMongoKey) ValidateOperater() interface{} {
	valid := validation.Validation{}
	b, _ := valid.Valid(v)
	if !b {
		// 表示获取验证的结构体
		st := reflect.TypeOf(ValidateMongoKey{})
		for _, err := range valid.Errors {
			// 获取验证的字段名和提示信息的别名
			filed, _ := st.FieldByName(err.Field)
			var alias = filed.Tag.Get("alias")
			//返回验证的错误信息
			errMsg := fmt.Sprintf("%s", errors.New(alias+err.Message))
			return errMsg
		}
	}
	return nil
}

// 验证是否有配置mongodb

// 验证是否连接的上mongodb
