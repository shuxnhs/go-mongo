package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

/**
 * 项目配置
 */

type Project struct {
	Id           int    `json:"id"`
	Project_name string `json:"project_name"`
	Mongo_key    string `json:"mongo_key"`
	Is_deleted   int    `json:"is_deleted"`
}

func AddProject(project Project) (string, error) {
	o := orm.NewOrm()
	MongoKey := generateMongoKey(project.Project_name)
	project.Mongo_key = MongoKey
	exist := o.QueryTable("project").Filter("Project_name", project.Project_name).Exist()
	if exist {
		return "", errors.New("exist")
	}
	_, err := o.Insert(&project)
	if err != nil {
		return "", errors.New("插入失败")
	}
	return MongoKey, nil
}

func GetAllProject() ([]*Project, error) {
	o := orm.NewOrm()
	var projects []*Project
	exist := o.QueryTable("Project").Exist()
	if exist {
		_, err := o.QueryTable("Project").All(&projects)
		if err != nil {
			return projects, nil
		}
		return nil, err
	}
	return nil, nil
}

// 生成唯一的mongo-key
func generateMongoKey(projectName string) string {
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+projectName))))
}
