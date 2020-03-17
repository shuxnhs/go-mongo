package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
)

type GoMongoSDK struct {
	GoMongoHost    string
	GoMongoVersion string
	GoMongoKey     string
}

// 返回格式
type APIResponse struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewGoMongoSDK() *GoMongoSDK {
	return &GoMongoSDK{
		GoMongoHost:    "",
		GoMongoVersion: "v1",
		GoMongoKey:     "",
	}
}

func (g *GoMongoSDK) SetMongoKey(goMongoKey string) *GoMongoSDK {
	g.GoMongoKey = goMongoKey
	return g
}

func (g *GoMongoSDK) SetMongoHost(goMongoHost string) *GoMongoSDK {
	g.GoMongoHost = goMongoHost
	return g
}

func (g *GoMongoSDK) Request(router string, service string, params map[string]string, reqType string) (APIResponse, error) {
	url := g.GoMongoHost + "/" + g.GoMongoVersion + "/" + router + "/" + service
	params["mongoKey"] = g.GoMongoKey
	return g.httpRequest(url, reqType, params)
}

func (g *GoMongoSDK) httpRequest(url string, reqType string, params map[string]string) (APIResponse, error) {
	apiResponse := &APIResponse{}
	if len(params) != 0 {
		url += "?"
		for key, value := range params {
			url += key + "=" + url2.QueryEscape(value) + "&"
		}
	}
	var resp *http.Response
	var err error
	if strings.ToUpper(reqType) == "GET" {
		resp, err = http.Get(url)
	} else {
		resp, err = http.Post(url, "", nil)
	}
	if err != nil {
		return *apiResponse, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	bytes, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(bytes, apiResponse)
	return *apiResponse, nil
}
