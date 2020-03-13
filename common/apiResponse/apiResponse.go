package apiResponse

import (
	"fmt"
)

// 接口返回的data部分数据
type APIDataResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewAPIDataResponse() APIDataResponse {
	return APIDataResponse{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}

func (a *APIDataResponse) Init() *APIDataResponse {
	a.Code = 0
	a.Msg = ""
	a.Data = nil
	return a
}

func (a *APIDataResponse) SetCode(code int) *APIDataResponse {
	a.Code = code
	return a
}

func (a *APIDataResponse) SetMsg(msg string) *APIDataResponse {
	a.Msg = msg
	return a
}

func (a *APIDataResponse) SetData(data interface{}) *APIDataResponse {
	a.Data = data
	return a
}

/** ------------------- 分割 ------------------- **/

// 顶层接口返回
type APIResponse struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 创建一个返回实例
func NewAPIResponse() APIResponse {
	data := APIDataResponse{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
	return APIResponse{
		Ret:  200,
		Msg:  "",
		Data: data,
	}
}

// 设置ret状态码
func (r *APIResponse) SetRet(ret int) *APIResponse {
	r.Ret = ret
	return r
}

// 设置（错误）提示信息 msg
func (r *APIResponse) SetMsg(msg string) *APIResponse {
	r.Msg = msg
	return r
}

// 设置返回的具体业务数据
func (r *APIResponse) SetData(data APIDataResponse) *APIResponse {
	r.Data = data
	return r
}

// 快捷接口，系统错误
func (r *APIResponse) Error(ret int, msg string) *APIResponse {
	r.SetRet(ret).SetMsg(msg)
	return r
}

// 快捷接口，业务失败
func (r *APIResponse) Fail(code int, msg string) *APIResponse {
	r.SetData(APIDataResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
	r.SetRet(200).SetMsg("")
	return r
}

// 快捷接口，业务成功返回
func (r *APIResponse) Success(data interface{}) *APIResponse {
	r.SetData(APIDataResponse{
		Code: 0,
		Msg:  "",
		Data: data,
	})
	r.SetRet(200).SetMsg("")
	return r
}

// 辅助方法
func (r *APIResponse) Print() {
	fmt.Printf("APIResponse: Ret = %d, Msg = %s \n", r.Ret, r.Msg)
}
