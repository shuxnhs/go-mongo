package validate

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"reflect"
)

//tags中 alias 表示验证不通过时候的提示名称，valid表示验证的格式 分号分割不同的验证内容
type ValidateProject struct {
	ProjectName string `alias:"项目名"  valid:"Required"`
}

// 验证添加配置的参数
func (v *ValidateProject) ValidAddProject() error {
	valid := validation.Validation{}
	b, _ := valid.Valid(v)
	if !b {
		//表示获取验证的结构体
		st := reflect.TypeOf(ValidateProject{})
		for _, err := range valid.Errors {
			//获取验证的字段名和提示信息的别名
			filed, _ := st.FieldByName(err.Field)
			var alias = filed.Tag.Get("alias")
			//返回验证的错误信息:fmt.Sprintf("%s", errors.New(alias+err.Message))
			return errors.New(alias + err.Message)
		}
	}
	return nil
}
