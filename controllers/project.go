package controllers

import (
	"fmt"
	"go-mongo/models"
	"go-mongo/validate"
)

// 项目相关接口
type ProjectController struct {
	BaseController
}

// @Title 添加新的项目
// @Description 添加新的项目，并返回分配的mongo-key
// @Param	ProjectName		query 	string	true		"项目名"
// @Success 0 {string} 添加成功
// @Failure 1 {string} 添加失败
// @router /AddProject [get]
func (ctx *ProjectController) AddProject() {
	// 参数获取
	pn := ctx.GetString("ProjectName")

	// 参数校验
	v := &validate.ValidateProject{
		ProjectName: pn,
	}
	vErr := v.ValidAddProject()
	if vErr != nil {
		ctx.ApiError(400, fmt.Sprintf("%s", vErr))
	}
	project := models.Project{
		Project_name: pn,
		Is_deleted:   0,
	}
	// 添加新项目
	mongoKey, err := models.AddProject(project)
	if err != nil {
		ctx.ApiFail(1, fmt.Sprintf("%s", err))
	}
	ctx.ApiSuccessData("添加成功", mongoKey)
}

/**----------------后台接口------------------**/

// @Title: 获取项目全部配置
func (ctx *ProjectController) GetAllProject() {
	projects, err := models.GetAllProject()
	if err != nil {
		ctx.ApiFail(1, fmt.Sprintf("%s", err))
	}
	ctx.ApiSuccessData("获取成功", projects)
}

// @Title: 后台管理界面首页
func (ctx *ProjectController) Index() {
	ctx.TplName = "index.html"
}

// @Title: 后台管理界面添加页
func (ctx *ProjectController) Add() {
	ctx.TplName = "add.html"
}
