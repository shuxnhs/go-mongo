package main

import "go-mongo/common/rsa"

// 用来生成公钥和私钥的小脚本。项目上线后禁止运行
func main() {
	rsa.GenerateRSAKey(1024)
}
