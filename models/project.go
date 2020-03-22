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

type projectInfo struct {
	Project
	Config
	// project和config都有MongoKey会冲突
	MongoKey string
}

func GetAllProject() ([]projectInfo, error) {
	o := orm.NewOrm()
	var projects []*Project
	_, err := o.QueryTable("project").All(&projects)
	if err != nil {
		return []projectInfo{}, nil
	}
	var projectInfos []projectInfo
	for i := 0; i < len(projects); i++ {
		mongoKey := projects[i].Mongo_key
		config, err := GetMongoConfig(mongoKey)
		var elem projectInfo
		if err != nil {
			elem = projectInfo{
				Project:  *projects[i],
				Config:   Config{},
				MongoKey: mongoKey,
			}
		}
		elem = projectInfo{
			Project:  *projects[i],
			Config:   config,
			MongoKey: mongoKey,
		}
		projectInfos = append(projectInfos, elem)
	}
	return projectInfos, err
}

// 生成唯一的mongo-key
func generateMongoKey(projectName string) string {
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String()+projectName))))
}
